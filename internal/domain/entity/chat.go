package entity

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	IsAnswer  bool      `json:"is_answer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	ContractID uuid.UUID `json:"contract_id"`
}
