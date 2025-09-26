package usecase

import (
	"bytes"
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/entity"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/enum"
	"github.com/Clausia-Ifest/clausia-server/pkg/pagination"
	clausiapb "github.com/Clausia-Ifest/clausia-server/proto"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (u *UContract) Update(ctx context.Context, req dto.UpdateContractRequest) error {
	tx, err := u.tx.Begin(ctx, false)
	if err != nil {
		return err
	}

	ap := enum.ParseAS(req.ApplicationStatus)
	if ap != enum.ASManager {
		return errors.New("application status not found")
	}

	st := enum.ParseStatus(req.Status)
	if st != enum.StatusRejected && st != enum.StatusAccepted {
		return errors.New("status not found")
	}

	_contract := &entity.Contract{
		ID:                req.ID,
		ApplicationStatus: ap,
		Status:            st,
		Notes:             req.Notes,
	}

	if err := u.rc.Update(ctx, tx.E, _contract); err != nil {
		return err
	}

	return nil
}

func (u *UContract) All(ctx context.Context, req dto.AllContractsRequest) (*dto.AllContractResponse, error) {
	tx, err := u.tx.Begin(ctx, false)
	if err != nil {
		return nil, err
	}

	params := dto.GetContractParams{
		Category:          req.Category,
		Status:            req.Status,
		ApplicationStatus: req.ApplicationStatus,
	}

	r, p, err := u.rc.All(ctx, tx.E, params)
	if err != nil {
		return nil, err
	}

	contracts := make([]*dto.Contract, len(r))
	for i, c := range r {
		contracts[i] = c.ParseDTO()
	}

	return &dto.AllContractResponse{
		Contracts:  contracts,
		Pagination: pagination.Initialize(p, req.Limit, req.Page),
	}, nil
}

func (u *UContract) Submit(ctx context.Context, req dto.SubmitContractRequest) error {
	tx, err := u.tx.Begin(ctx, true)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			return
		}
	}()

	category := enum.ParseCategory(req.Category)
	if category == enum.CategoryUnknown {
		return errors.New("category not found")
	}

	newContract := entity.Contract{
		ID:        uuid.New(),
		HumanID:   randomString(15),
		Email:     req.Email,
		Title:     req.Title,
		Company:   req.Company,
		Category:  category,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}
	if err := u.rc.Create(ctx, tx.E, newContract); err != nil {
		return err
	}

	for c, document := range []*multipart.FileHeader{
		req.AdministrationDocument,
		req.LegalDocument,
		req.TechnicalDocument,
		req.FinancialDocument,
	} {
		var (
			params dto.GetDocumentParams
		)

		hash, err := hashSHA256(document)
		if err != nil {
			return err
		}

		params.Hash = hash

		src, err := document.Open()
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		h := sha256.New()
		if _, err := io.Copy(io.MultiWriter(&buf, h), src); err != nil {
			src.Close()
			return err
		}
		src.Close()

		if buf.Len() < 4 || !bytes.Equal(buf.Bytes()[:4], []byte("%PDF")) {
			return errors.New("only pdf files are allowed")
		}

		reader := bytes.NewReader(buf.Bytes())

		_document, err := u.rd.Get(ctx, tx.E, params)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		if errors.Is(err, sql.ErrNoRows) {
			var contractDocument entity.ContractDocument

			ctxNoCancel := context.WithoutCancel(context.Background())
			go func() {
				notx, err := u.tx.Begin(ctxNoCancel, false)
				if err != nil {
					log.Err(err).Str("hash", hash).Msg("failed to begin tx")
					return
				}

				if _, err := reader.Seek(0, io.SeekStart); err != nil {
					log.Err(err).Str("hash", hash).Msg("failed to seek file")
					return
				}

				if err := u.s3.Upload(ctxNoCancel, "documents/"+hash, reader, "application/pdf"); err != nil {
					log.Err(err).Str("hash", hash).Msg("failed upload to s3")
					return
				}

				p, err := u.grpc.ExtractMetadata(ctxNoCancel, &clausiapb.ExtractRequest{
					Source: &clausiapb.ExtractRequest_S3Ref{
						S3Ref: &clausiapb.S3Reference{
							ObjectKey: hash,
						},
					},
				})
				if err != nil {
					log.Err(err).Str("hash", hash).Msg("failed to hit ai")
					return
				}

				_document = &entity.Document{
					Hash:     hash,
					MetaData: p.Metadata,
					Content:  p.Content,
				}

				if err := u.rd.Create(ctxNoCancel, notx.E, *_document); err != nil {
					log.Err(err).Str("hash", hash).Msg("failed store to db")
					return
				}

				c := enum.DocumentCategory(c)
				if c.String() == "" {
					log.Err(errors.New("category not found")).Str("hash", hash).Msg("category not found")
					return
				}

				contractDocument = entity.ContractDocument{
					DocumentHash: hash,
					ContractID:   newContract.ID,
					URL:          fmt.Sprintf("documents/%s", hash),
					Category:     c,
				}

				if err := u.rc.CreateRelation(ctxNoCancel, notx.E, contractDocument); err != nil {
					log.Err(err).Str("hash", hash).Msg("failed store to db - relation")
					return
				}
			}()
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
