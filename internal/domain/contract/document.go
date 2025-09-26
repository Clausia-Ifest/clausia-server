package contract

import (
	"context"
	"mime/multipart"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type IRDocument interface {
	Get(ctx context.Context, ext sqlx.ExtContext, conditions dto.GetDocumentParams) (*entity.Document, error)
	Create(ctx context.Context, ext sqlx.ExtContext, data entity.Document) error
}

type IUDocument interface {
	Extract(ctx context.Context, document *multipart.FileHeader) (*dto.ExtractDocumentResponse, error)
}
