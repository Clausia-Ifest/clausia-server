package db

import (
	"errors"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

func (db *Database) Migrate(command string) error {
	driver, err := postgres.WithInstance(db.Conn.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	switch command {
	case "up":
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Info().
					Msg("[MIGRATE][UP] No change - schema is up to date")

			} else {
				log.Error().
					Err(err).
					Msg("[MIGRATE][UP] Failed to apply UP migrations")

				os.Exit(0)
			}
		}

		log.Info().
			Msg("[MIGRATE][UP] Applied UP migrations successfully")
	case "down":
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Info().
					Msg("[MIGRATE][DOWN] No change - nothing to roll back")

			} else {
				log.Error().
					Err(err).
					Msg("[MIGRATE][DOWN] Failed to apply DOWN migrations")

				os.Exit(0)
			}
		}

		log.Info().
			Msg("[MIGRATE][DOWN] Applied DOWN migrations successfully")
	case "fresh":
		if err := m.Drop(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Info().
					Msg("[MIGRATE][FRESH] Already at version 0 - nothing to reset")

			} else {
				log.Error().
					Err(err).
					Msg("[MIGRATE][FRESH] Failed to reset schema to version 0")

				os.Exit(0)
			}
		}

		// https://github.com/golang-migrate/migrate/issues/226
		driver, err = postgres.WithInstance(db.Conn.DB, &postgres.Config{})
		if err != nil {
			return err
		}

		m, err = migrate.NewWithDatabaseInstance(
			"file://db/migrations",
			"postgres", driver)
		if err != nil {
			return err
		}

		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Info().
					Msg("[MIGRATE][FRESH] No change - schema is up to date")

			} else {
				log.Error().
					Err(err).
					Msg("[MIGRATE][FRESH] Failed to apply UP migrations")

				os.Exit(0)
			}
		}

		log.Info().
			Msg("[MIGRATE][FRESH] Database refreshed successfully")
	}

	return nil
}
