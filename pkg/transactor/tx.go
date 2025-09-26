package transactor

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Transactor struct {
	db *sqlx.DB
}

type ITransactor interface {
	Begin(ctx context.Context, withTx bool) (*Handle, error)
}

type Handle struct {
	E  sqlx.ExtContext
	tx *sqlx.Tx
}

func New(db *sqlx.DB) ITransactor {
	return &Transactor{
		db: db,
	}
}

func (t *Transactor) Begin(ctx context.Context, withTx bool) (*Handle, error) {
	if withTx {
		tx, err := t.db.BeginTxx(ctx, nil)
		if err != nil {
			return nil, err
		}

		return &Handle{
			E:  tx,
			tx: tx,
		}, nil
	} else {
		return &Handle{
			E:  t.db,
			tx: nil,
		}, nil
	}
}

func (h *Handle) Commit() error {
	if h.tx != nil {
		return h.tx.Commit()
	}

	return errors.New("failed to commit")
}

func (h *Handle) Rollback() error {
	if h.tx != nil {
		return h.tx.Rollback()
	}

	return errors.New("failed to rollback")
}
