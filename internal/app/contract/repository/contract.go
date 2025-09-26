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

func (r *RContract) Update(ctx context.Context, ext sqlx.ExtContext, data *entity.Contract) error {
	query := `
	UPDATE 
		contracts
	SET 
		application_status = :application_status,
		status = :status,
		notes = :notes
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
			c.id AS contract_id,
			c.human_id AS contract_human_id,
			c.email AS contract_email,
			c.title AS contract_title,
			c.company AS contract_company,
			c.notes AS contract_notes,
			c.risk_level AS contract_risk_level,
			c.status AS contract_status,
			c.application_status AS contract_application_status,
			c.category AS contract_category,
			c.start_date AS contract_start_date,
			c.end_date AS contract_end_date,
			c.created_at AS contract_created_at,
			c.updated_at AS contract_updated_at,
			COUNT(*) OVER() AS total_count
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
		cd.document_hash      AS contract_doc_hash,
		cd.url      AS contract_doc_url,
		cd.category AS contract_doc_category
	FROM base b
	LEFT JOIN contract_documents cd
		ON cd.contract_id = b.contract_id
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
	}

	if err := sqlx.SelectContext(ctx, ext, &rows, query); err != nil {
		return nil, 0, err
	}
	if len(rows) == 0 {
		return nil, 0, sql.ErrNoRows
	}

	byID := make(map[uuid.UUID]*entity.Contract, len(rows))
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
			total = rrow.TotalCount
			ctr = tmp
		}

		if rrow.ContractDocHash != nil || rrow.ContractDocURL != nil || rrow.ContractDocCategory != nil {
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
			ctr.ContractDocument = append(ctr.ContractDocument, doc)
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
