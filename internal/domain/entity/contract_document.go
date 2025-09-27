package entity

import (
	"time"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/enum"
	"github.com/google/uuid"
)

type ContractDocument struct {
	DocumentHash string    `db:"document_hash"`
	ContractID   uuid.UUID `db:"contract_id"`

	Content   string                `db:"content"`
	MetaData  string                `db:"meta_data"`
	URL       string                `db:"url"`
	Category  enum.DocumentCategory `db:"category"`
	CreatedAt time.Time             `db:"created_at"`
	UpdatedAt time.Time             `db:"updated_at"`
}
