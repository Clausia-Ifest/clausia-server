package dto

import (
	"mime/multipart"
	"time"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/enum"
	"github.com/Clausia-Ifest/clausia-server/pkg/pagination"
	"github.com/google/uuid"
)

type Contract struct {
	ID                uuid.UUID `json:"id"`
	HumanID           string    `json:"human_id"`
	Email             string    `json:"email"`
	Title             string    `json:"title"`
	Company           string    `json:"company"`
	Notes             string    `json:"notes"`
	RiskDetection     string    `json:"risk_detection"`
	Summarize         string    `json:"summarize"`
	RiskLevel         string    `json:"risk_level"`
	Status            string    `json:"status"`
	ApplicationStatus string    `json:"application_status"`
	Category          string    `json:"category"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`

	Documents []Document `json:"documents"`
	Chats     []Chat     `json:"chats"`
}

type SubmitContractRequest struct {
	Email     string    `form:"email"`
	Title     string    `form:"title"`
	Company   string    `form:"company"`
	Category  string    `form:"category"`
	StartDate time.Time `form:"start_date"`
	EndDate   time.Time `form:"end_date"`

	AdministrationDocument *multipart.FileHeader `form:"administration_document"`
	LegalDocument          *multipart.FileHeader `form:"legal_document"`
	TechnicalDocument      *multipart.FileHeader `form:"technical_document"`
	FinancialDocument      *multipart.FileHeader `form:"financial_document"`
}

type AllContractsRequest struct {
	Limit             int64
	Page              int64
	StartDate         time.Time
	EndDate           time.Time
	Category          enum.Category
	Status            enum.Status
	ApplicationStatus enum.ApplicationStatus
}

type AllContractResponse struct {
	Contracts  []*Contract                   `json:"contracts"`
	Pagination pagination.PaginationMetaData `json:"pagination"`
}

type GetContractParams struct {
	ID                uuid.UUID              `db:"id"`
	StartDate         time.Time              `db:"-"`
	EndDate           time.Time              `db:"-"`
	Category          enum.Category          `db:"category"`
	Status            enum.Status            `db:"status"`
	ApplicationStatus enum.ApplicationStatus `db:"application_status"`
}

type UpdateContractRequest struct {
	ID                uuid.UUID
	ApplicationStatus string `json:"application_status"`
	Status            string `json:"status"`
	Notes             string `json:"notes"`
	UserRole          string `json:"-"`
}
