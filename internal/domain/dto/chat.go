package dto

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	IsAnswer  bool      `json:"is_answer"`
	CreatedAt time.Time `json:"created_at"`
}

type ChatRequest struct {
	ContractID uuid.UUID
	Message    string `json:"message"`
}
