package contract

import (
	"context"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IRContract interface {
	Create(ctx context.Context, ext sqlx.ExtContext, data entity.Contract) error
	CreateRelation(ctx context.Context, ext sqlx.ExtContext, data entity.ContractDocument) error
	All(ctx context.Context, ext sqlx.ExtContext, conditions dto.GetContractParams) ([]entity.Contract, int64, error)
	Update(ctx context.Context, ext sqlx.ExtContext, data *entity.Contract) error
	Get(ctx context.Context, ext sqlx.ExtContext, conditions dto.GetContractParams) (*entity.Contract, error)
	CreateChatHistory(ctx context.Context, ext sqlx.ExtContext, data entity.Chat) error
}

type IUContract interface {
	Submit(ctx context.Context, req dto.SubmitContractRequest) error
	All(ctx context.Context, req dto.AllContractsRequest) (*dto.AllContractResponse, error)
	Update(ctx context.Context, req dto.UpdateContractRequest) error
	Get(ctx context.Context, id uuid.UUID) (*dto.Contract, error)
	Chat(ctx context.Context, req dto.ChatRequest) (string, error)
}
