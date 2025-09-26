package entity

import (
	"time"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/enum"
	"github.com/google/uuid"
)

type Contract struct {
	ID                uuid.UUID              `db:"id"`
	HumanID           string                 `db:"human_id"`
	Email             string                 `db:"email"`
	Title             string                 `db:"title"`
	Company           string                 `db:"company"`
	Notes             string                 `db:"notes"`
	RiskLevel         enum.RiskLevel         `db:"risk_level"`
	Status            enum.Status            `db:"status"`
	ApplicationStatus enum.ApplicationStatus `db:"application_status"`
	Category          enum.Category          `db:"category"`
	StartDate         time.Time              `db:"start_date"`
	EndDate           time.Time              `db:"end_date"`
	CreatedAt         time.Time              `db:"created_at"`
	UpdatedAt         time.Time              `db:"updated_at"`

	ContractDocument []ContractDocument `db:"contract_document"`
}

func (e *Contract) ParseDTO() *dto.Contract {
	r := &dto.Contract{
		ID:                e.ID,
		HumanID:           e.HumanID,
		Email:             e.Email,
		Title:             e.Title,
		Company:           e.Company,
		Notes:             e.Notes,
		RiskLevel:         e.RiskLevel.String(),
		Status:            e.Status.String(),
		ApplicationStatus: e.ApplicationStatus.String(),
		Category:          e.Category.String(),
		StartDate:         e.StartDate,
		EndDate:           e.EndDate,

		Documents: make([]dto.Document, 0, len(e.ContractDocument)),
	}

	for _, d := range e.ContractDocument {
		r.Documents = append(r.Documents, dto.Document{
			Hash:     d.DocumentHash,
			URL:      d.URL,
			Category: d.Category.String(),
		})
	}

	return r
}
