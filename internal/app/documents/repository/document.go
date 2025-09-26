package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

func (r *RDocument) Get(ctx context.Context, ext sqlx.ExtContext, conditions dto.GetDocumentParams) (*entity.Document, error) {
	query := `
	SELECT
		d.hash,
		d.meta_data,
		d.content,
		d.created_at
	FROM
		documents d
	`

	c := r.qb.WhereConditions(conditions)
	if c != "" {
		query += fmt.Sprintf(`WHERE %s`, c)
	}

	row, err := sqlx.NamedQueryContext(ctx, ext, query, conditions)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	e := new(entity.Document)

	if row.Next() {
		if err := row.StructScan(e); err != nil {
			return nil, err
		}

		return e, nil
	}

	return nil, sql.ErrNoRows
}

func (r *RDocument) Create(ctx context.Context, ext sqlx.ExtContext, data entity.Document) error {
	query := `
	INSERT INTO documents (
		hash,
		meta_data,
		content
	) VALUES (
	 	:hash,
		:meta_data,
		:content
	);
	`

	_, err := sqlx.NamedExecContext(ctx, ext, query, data)
	return err
}
