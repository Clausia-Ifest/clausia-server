package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/entity"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/enum"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (r *RContract) Get(ctx context.Context, ext sqlx.ExtContext, conditions dto.GetContractParams) (*entity.Contract, error) {
	where := r.qb.WhereConditions(conditions)
	if where != "" {
		where = "WHERE " + where
	}

	query := fmt.Sprintf(`
		WITH base AS (
			SELECT
				c.id                  AS contract_id,
				c.human_id            AS contract_human_id,
				c.email               AS contract_email,
				c.title               AS contract_title,
				c.company             AS contract_company,
				c.notes               AS contract_notes,
				c.risk_level          AS contract_risk_level,
				c.status              AS contract_status,
				c.application_status  AS contract_application_status,
				c.category            AS contract_category,
				c.start_date          AS contract_start_date,
				c.end_date            AS contract_end_date,
				c.created_at          AS contract_created_at,
				c.updated_at          AS contract_updated_at,
				c.risk_detection      AS contract_risk_detection,
				c.summarize           AS contract_summarize
			FROM contracts c
			%s
			ORDER BY c.created_at DESC
			LIMIT 1
		)
		SELECT
			b.*,
			TRIM(BOTH FROM cd.document_hash) AS contract_doc_hash, -- hilangin padding CHAR(64)
			d.content                        AS contract_doc_content,
			cd.url                           AS contract_doc_url,
			cd.category                      AS contract_doc_category,
			d.meta_data                   AS contract_doc_metadata, -- aktifkan kalau kamu pakai
			ch.id                            AS chat_id,
			ch.content                       AS chat_content,
			ch.is_answer                     AS chat_is_answer,
			ch.created_at                    AS chat_created_at,
			ch.updated_at                    AS chat_updated_at
		FROM base b
		LEFT JOIN contract_documents cd
			ON cd.contract_id = b.contract_id
		LEFT JOIN documents d
			ON d.hash = cd.document_hash
		LEFT JOIN chats ch
			ON ch.contract_id = b.contract_id
	`, where)

	type row struct {
		ContractID                uuid.UUID `db:"contract_id"`
		ContractHumanID           string    `db:"contract_human_id"`
		ContractEmail             string    `db:"contract_email"`
		ContractTitle             string    `db:"contract_title"`
		ContractCompany           string    `db:"contract_company"`
		ContractNotes             *string   `db:"contract_notes"`
		ContractRiskLevel         int64     `db:"contract_risk_level"`
		ContractStatus            int64     `db:"contract_status"`
		ContractApplicationStatus int64     `db:"contract_application_status"`
		ContractCategory          int64     `db:"contract_category"`
		ContractStartDate         time.Time `db:"contract_start_date"`
		ContractEndDate           time.Time `db:"contract_end_date"`
		ContractCreatedAt         time.Time `db:"contract_created_at"`
		ContractUpdatedAt         time.Time `db:"contract_updated_at"`
		ContractRiskDetection     *string   `db:"contract_risk_detection"`
		ContractSummarize         *string   `db:"contract_summarize"`

		ContractDocHash     *string `db:"contract_doc_hash"`
		ContractDocContent  *string `db:"contract_doc_content"`
		ContractDocURL      *string `db:"contract_doc_url"`
		ContractDocCategory *int64  `db:"contract_doc_category"`
		ContractDocMetadata *string `db:"contract_doc_metadata"`

		ChatID        *uuid.UUID `db:"chat_id"`
		ChatContent   *string    `db:"chat_content"`
		ChatIsAnswer  *bool      `db:"chat_is_answer"`
		ChatCreatedAt *time.Time `db:"chat_created_at"`
		ChatUpdatedAt *time.Time `db:"chat_updated_at"`
	}

	rowsx, err := sqlx.NamedQueryContext(ctx, ext, query, conditions)
	if err != nil {
		return nil, err
	}
	defer rowsx.Close()

	var rows []row
	for rowsx.Next() {
		var r row
		if err := rowsx.StructScan(&r); err != nil {
			return nil, err
		}
		rows = append(rows, r)
	}
	if err := rowsx.Err(); err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, sql.ErrNoRows
	}

	first := rows[0]
	ctr := &entity.Contract{
		ID:                first.ContractID,
		HumanID:           first.ContractHumanID,
		Email:             first.ContractEmail,
		Title:             first.ContractTitle,
		Company:           first.ContractCompany,
		RiskLevel:         enum.RiskLevel(first.ContractRiskLevel),
		Status:            enum.Status(first.ContractStatus),
		ApplicationStatus: enum.ApplicationStatus(first.ContractApplicationStatus),
		Category:          enum.Category(first.ContractCategory),
		StartDate:         first.ContractStartDate,
		EndDate:           first.ContractEndDate,
		CreatedAt:         first.ContractCreatedAt,
		UpdatedAt:         first.ContractUpdatedAt,
		ContractDocument:  make([]entity.ContractDocument, 0),
		ContractChats:     make([]entity.Chat, 0),
	}

	if first.ContractNotes != nil {
		ctr.Notes = *first.ContractNotes
	}
	if first.ContractRiskDetection != nil {
		ctr.RiskDetection = *first.ContractRiskDetection
	}
	if first.ContractSummarize != nil {
		ctr.Summarize = *first.ContractSummarize
	}

	// De-dupe untuk menghindari duplikasi akibat join (doc × chat)
	seenDocs := make(map[string]struct{})
	seenChats := make(map[uuid.UUID]struct{})

	for _, rr := range rows {
		// Document
		if rr.ContractDocHash != nil || rr.ContractDocContent != nil || rr.ContractDocMetadata != nil || rr.ContractDocURL != nil || rr.ContractDocCategory != nil {
			key := ""
			if rr.ContractDocHash != nil {
				key = *rr.ContractDocHash
			} else if rr.ContractDocURL != nil {
				key = *rr.ContractDocURL // fallback
			}
			if _, ok := seenDocs[key]; !ok {
				doc := entity.ContractDocument{}
				if rr.ContractDocContent != nil {
					doc.Content = *rr.ContractDocContent
				}
				if rr.ContractDocHash != nil {
					doc.DocumentHash = *rr.ContractDocHash
				}
				if rr.ContractDocURL != nil {
					doc.URL = *rr.ContractDocURL
				}
				if rr.ContractDocCategory != nil {
					doc.Category = enum.DocumentCategory(*rr.ContractDocCategory)
				}
				if rr.ContractDocMetadata != nil {
					doc.MetaData = *rr.ContractDocMetadata
				}

				ctr.ContractDocument = append(ctr.ContractDocument, doc)
				if key != "" {
					seenDocs[key] = struct{}{}
				}
			}
		}

		// Chat
		if rr.ChatID != nil && rr.ChatContent != nil {
			if _, ok := seenChats[*rr.ChatID]; !ok {
				chat := entity.Chat{
					ID:        *rr.ChatID,
					Content:   *rr.ChatContent,
					IsAnswer:  false,
					CreatedAt: *rr.ChatCreatedAt,
					UpdatedAt: *rr.ChatUpdatedAt,
				}
				if rr.ChatIsAnswer != nil {
					chat.IsAnswer = *rr.ChatIsAnswer
				}
				ctr.ContractChats = append(ctr.ContractChats, chat)
				seenChats[*rr.ChatID] = struct{}{}
			}
		}
	}

	return ctr, nil
}

func (r *RContract) Update(ctx context.Context, ext sqlx.ExtContext, data *entity.Contract) error {
	query := `
	UPDATE 
		contracts
	SET 
		application_status = :application_status,
		status = :status,
		notes = :notes,
		risk_detection = :risk_detection,
		risk_level = :risk_level
	WHERE
		id = :id;
	`

	_, err := sqlx.NamedExecContext(ctx, ext, query, data)
	return err
}

func (r *RContract) All(ctx context.Context, ext sqlx.ExtContext, conditions dto.GetContractParams) ([]entity.Contract, int64, error) {
	where := r.qb.WhereConditions(conditions)
	if where != "" {
		where = "WHERE " + where
	}

	query := fmt.Sprintf(`
	WITH base AS (
		SELECT
			c.id                  AS contract_id,
			c.human_id            AS contract_human_id,
			c.email               AS contract_email,
			c.title               AS contract_title,
			c.company             AS contract_company,
			c.notes               AS contract_notes,
			c.risk_level          AS contract_risk_level,
			c.status              AS contract_status,
			c.application_status  AS contract_application_status,
			c.category            AS contract_category,
			c.start_date          AS contract_start_date,
			c.end_date            AS contract_end_date,
			c.created_at          AS contract_created_at,
			c.updated_at          AS contract_updated_at,
			COUNT(*) OVER()       AS total_count
		FROM contracts c
		%s
	)
	SELECT
		b.contract_id,
		b.contract_human_id,
		b.contract_email,
		b.contract_title,
		b.contract_company,
		b.contract_notes,
		b.contract_risk_level,
		b.contract_status,
		b.contract_application_status,
		b.contract_category,
		b.contract_start_date,
		b.contract_end_date,
		b.contract_created_at,
		b.contract_updated_at,
		b.total_count,
		TRIM(BOTH FROM cd.document_hash) AS contract_doc_hash,
		cd.url                           AS contract_doc_url,
		cd.category                      AS contract_doc_category,
		d.content                        AS contract_doc_content,
		d.meta_data                      AS contract_doc_metadata
	FROM base b
	LEFT JOIN contract_documents cd
		ON cd.contract_id = b.contract_id
	LEFT JOIN documents d
		ON TRIM(BOTH FROM d.hash) = TRIM(BOTH FROM cd.document_hash)
	`, where)

	var rows []struct {
		ContractID                uuid.UUID `db:"contract_id"`
		ContractHumanID           string    `db:"contract_human_id"`
		ContractEmail             string    `db:"contract_email"`
		ContractTitle             string    `db:"contract_title"`
		ContractCompany           string    `db:"contract_company"`
		ContractNotes             *string   `db:"contract_notes"`
		ContractRiskLevel         int64     `db:"contract_risk_level"`
		ContractStatus            int64     `db:"contract_status"`
		ContractApplicationStatus int64     `db:"contract_application_status"`
		ContractCategory          int64     `db:"contract_category"`
		ContractStartDate         time.Time `db:"contract_start_date"`
		ContractEndDate           time.Time `db:"contract_end_date"`
		ContractCreatedAt         time.Time `db:"contract_created_at"`
		ContractUpdatedAt         time.Time `db:"contract_updated_at"`
		TotalCount                int64     `db:"total_count"`

		ContractDocHash     *string `db:"contract_doc_hash"`
		ContractDocURL      *string `db:"contract_doc_url"`
		ContractDocCategory *int64  `db:"contract_doc_category"`
		ContractDocContent  *string `db:"contract_doc_content"`
		ContractDocMetadata *string `db:"contract_doc_metadata"`
	}

	if err := sqlx.SelectContext(ctx, ext, &rows, query); err != nil {
		return nil, 0, err
	}
	if len(rows) == 0 {
		return nil, 0, sql.ErrNoRows
	}

	byID := make(map[uuid.UUID]*entity.Contract, len(rows))
	seenDocByContract := make(map[uuid.UUID]map[string]struct{}) // de-dupe doc per kontrak
	var total int64

	for _, rrow := range rows {
		ctr, ok := byID[rrow.ContractID]
		if !ok {
			tmp := &entity.Contract{
				ID:                rrow.ContractID,
				HumanID:           rrow.ContractHumanID,
				Email:             rrow.ContractEmail,
				Title:             rrow.ContractTitle,
				Company:           rrow.ContractCompany,
				RiskLevel:         enum.RiskLevel(rrow.ContractRiskLevel),
				Status:            enum.Status(rrow.ContractStatus),
				ApplicationStatus: enum.ApplicationStatus(rrow.ContractApplicationStatus),
				Category:          enum.Category(rrow.ContractCategory),
				StartDate:         rrow.ContractStartDate,
				EndDate:           rrow.ContractEndDate,
				CreatedAt:         rrow.ContractCreatedAt,
				UpdatedAt:         rrow.ContractUpdatedAt,
				ContractDocument:  make([]entity.ContractDocument, 0),
			}
			if rrow.ContractNotes != nil {
				tmp.Notes = *rrow.ContractNotes
			}
			byID[rrow.ContractID] = tmp
			seenDocByContract[rrow.ContractID] = make(map[string]struct{})
			total = rrow.TotalCount
			ctr = tmp
		}

		// append dokumen (de-dupe by hash/url)
		if rrow.ContractDocHash != nil || rrow.ContractDocURL != nil || rrow.ContractDocCategory != nil || rrow.ContractDocContent != nil || rrow.ContractDocMetadata != nil {
			key := ""
			if rrow.ContractDocHash != nil && *rrow.ContractDocHash != "" {
				key = *rrow.ContractDocHash
			} else if rrow.ContractDocURL != nil && *rrow.ContractDocURL != "" {
				key = *rrow.ContractDocURL
			}

			if _, exists := seenDocByContract[rrow.ContractID][key]; !exists {
				doc := entity.ContractDocument{}
				if rrow.ContractDocHash != nil {
					doc.DocumentHash = *rrow.ContractDocHash
				}
				if rrow.ContractDocURL != nil {
					doc.URL = *rrow.ContractDocURL
				}
				if rrow.ContractDocCategory != nil {
					doc.Category = enum.DocumentCategory(*rrow.ContractDocCategory)
				}
				if rrow.ContractDocContent != nil {
					doc.Content = *rrow.ContractDocContent
				}
				if rrow.ContractDocMetadata != nil {
					doc.MetaData = *rrow.ContractDocMetadata
				}

				ctr.ContractDocument = append(ctr.ContractDocument, doc)
				if key != "" {
					seenDocByContract[rrow.ContractID][key] = struct{}{}
				}
			}
		}
	}

	contracts := make([]entity.Contract, 0, len(byID))
	for _, c := range byID {
		contracts = append(contracts, *c)
	}

	return contracts, total, nil
}

func (r *RContract) Create(ctx context.Context, ext sqlx.ExtContext, data entity.Contract) error {
	query := `
	INSERT INTO contracts (
		id,
		human_id,
		email,
		title,
		company,
		risk_level,
		status,
		application_status,
		category,
		start_date,
		end_date
	) VALUES (
		:id,
		:human_id,
		:email,
		:title,
		:company,
		:risk_level,
		:status,
		:application_status,
		:category,
		:start_date,
		:end_date
	);
	`

	_, err := sqlx.NamedExecContext(ctx, ext, query, data)
	return err
}

func (r *RContract) CreateRelation(ctx context.Context, ext sqlx.ExtContext, data entity.ContractDocument) error {
	query := `
	INSERT INTO contract_documents (
		document_hash,
		contract_id,
		url,
		category
	) VALUES (
		:document_hash,
		:contract_id,
		:url,
		:category
	);
	`

	_, err := sqlx.NamedExecContext(ctx, ext, query, data)
	return err
}

func (r *RContract) CreateChatHistory(ctx context.Context, ext sqlx.ExtContext, data entity.Chat) error {
	query := `
	INSERT INTO chats (
		id,
		content,
		is_answer,
		contract_id
	) VALUES (
		:id,
		:content,
		:is_answer,
		:contract_id
	);
	`

	_, err := sqlx.NamedExecContext(ctx, ext, query, data)
	return err
}
