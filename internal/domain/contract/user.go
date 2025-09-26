package contract

import (
	"context"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IRUser interface {
	Get(ctx context.Context, ext sqlx.ExtContext, conditions dto.GetUserParams) (*entity.User, error)
}

type IUUser interface {
	Auth(ctx context.Context, req dto.SignInRequest) (*dto.SignInResponse, error)
	Self(ctx context.Context, userID uuid.UUID) (*dto.User, error)
}
