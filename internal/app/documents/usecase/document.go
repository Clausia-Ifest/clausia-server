package usecase

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"mime/multipart"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/entity"
	clausiapb "github.com/Clausia-Ifest/clausia-server/proto"
	"github.com/rs/zerolog/log"
)

func (u *UDocument) Extract(ctx context.Context, document *multipart.FileHeader) (*dto.ExtractDocumentResponse, error) {
	if err := assertPDF(document); err != nil {
		return nil, err
	}

	tx, err := u.tx.Begin(ctx, false)
	if err != nil {
		return nil, err
	}

	hash, err := hashSHA256(document)
	if err != nil {
		return nil, err
	}

	params := dto.GetDocumentParams{
		Hash: hash,
	}

	_document, err := u.rd.Get(ctx, tx.E, params)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		f, err := document.Open()
		if err != nil {
			return nil, err
		}
		defer f.Close()

		if _, err := f.Seek(0, io.SeekStart); err != nil {
			return nil, err
		}

		if err := u.s3.Upload(ctx, "documents/"+hash, f, "application/pdf"); err != nil {
			log.Err(err).Str("hash", hash).Msg("failed upload to s3")
			return nil, err
		}

		p, err := u.grpc.ExtractMetadata(ctx, &clausiapb.ExtractRequest{
			Source: &clausiapb.ExtractRequest_S3Ref{
				S3Ref: &clausiapb.S3Reference{
					ObjectKey: hash,
				},
			},
		})
		if err != nil {
			return nil, err
		}

		_document = &entity.Document{
			Hash:     hash,
			MetaData: p.Metadata,
			Content:  p.Content,
		}

		if err := u.rd.Create(ctx, tx.E, *_document); err != nil {
			return nil, err
		}
	}

	return &dto.ExtractDocumentResponse{
		MetaData: _document.MetaData,
	}, nil
}
