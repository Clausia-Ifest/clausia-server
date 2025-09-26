package db

import (
	"fmt"

	"github.com/Clausia-Ifest/clausia-server/internal/infra/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sqlx.DB
}

type Postgres interface {
	GetConnection() *sqlx.DB
	Migrate(command string) error
	Seed()
}

func NewPostgres(_config config.DB) (Postgres, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=public",
		_config.DatabaseHost,
		_config.DatabasePort,
		_config.DatabaseUsername,
		_config.DatabasePassword,
		_config.DatabaseDB,
		_config.DatabaseSSL,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return &Database{}, err
	}

	return &Database{
		Conn: db,
	}, nil
}

func (db *Database) GetConnection() *sqlx.DB {
	return db.Conn
}
