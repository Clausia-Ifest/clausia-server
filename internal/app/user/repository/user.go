package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

func (r *RUser) Get(ctx context.Context, ext sqlx.ExtContext, conditions dto.GetUserParams) (*entity.User, error) {
	query := `
	SELECT
		u.id,
		u.full_name,
		u.email,
		u.password,
		u.role,
		u.created_at,
		u.updated_at
	FROM
		users u 
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

	e := new(entity.User)

	if row.Next() {
		if err := row.StructScan(e); err != nil {
			return nil, err
		}

		return e, nil
	}

	return nil, sql.ErrNoRows
}
