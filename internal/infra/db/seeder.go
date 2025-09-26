package db

import (
	"fmt"

	"github.com/Clausia-Ifest/clausia-server/internal/infra/db/seeder"
	"github.com/rs/zerolog/log"
)

func (db *Database) Seed() {
	var successCount, errorCount int

	queryInsertUser := `
		INSERT INTO users (
			id,
			full_name,
			email,
			password,
			role
		) VALUES (
			:id,
			:full_name,
			:email,
			:password,
			:role
		);
	`

	successCount = 0
	errorCount = 0

	for _, p := range seeder.Users {
		_, err := db.Conn.NamedExec(queryInsertUser, p)
		if err != nil {
			log.Error().
				Err(err).
				Str("user_email", p.Email).
				Msg("[SEEDER][USER] Failed to insert user")

			errorCount++
			continue
		}

		successCount++
	}

	log.Info().
		Int("success", successCount).
		Int("errors", errorCount).
		Msg("[SEEDER][USER] Seeding completed")

	queryInsertDocument := `
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

	successCount = 0
	errorCount = 0

	for _, p := range seeder.Documents {
		_, err := db.Conn.NamedExec(queryInsertDocument, p)
		if err != nil {
			log.Error().
				Err(err).
				Str("document_hash", p.Hash).
				Msg("[SEEDER][DOCUMENT] Failed to insert document")

			errorCount++
			continue
		}

		successCount++
	}

	log.Info().
		Int("success", successCount).
		Int("errors", errorCount).
		Msg("[SEEDER][DOCUMENT] Seeding completed")

	queryInsertContract := `
		INSERT INTO contracts (
		id,
		human_id,
		email,
		title,
		company,
		risk_level,
		status,
		application_status,
		category,
		start_date,
		end_date
	) VALUES (
		:id,
		:human_id,
		:email,
		:title,
		:company,
		:risk_level,
		:status,
		:application_status,
		:category,
		:start_date,
		:end_date
	);
	`

	successCount = 0
	errorCount = 0

	for i, p := range seeder.Contracts {
		p.HumanID += fmt.Sprintf("%04d", i+1)

		_, err := db.Conn.NamedExec(queryInsertContract, p)
		if err != nil {
			log.Error().
				Err(err).
				Str("contract_title", p.Title).
				Msg("[SEEDER][CONTRACT] Failed to insert contract")

			errorCount++
			continue
		}

		for _, _p := range p.ContractDocument {
			queryInsertContractDocuments := `
			INSERT INTO contract_documents (
				document_hash,
				contract_id,
				url,
				category
			) VALUES (
				:document_hash,
				:contract_id,
				:url,
				:category
			);
			`

			_p.ContractID = p.ID
			_p.URL = fmt.Sprintf("https://is3.cloudhost.id/sembilan-belas-derajat-won-ifest-2025/documents/%s", _p.DocumentHash)

			_, err := db.Conn.NamedExec(queryInsertContractDocuments, _p)
			if err != nil {
				log.Error().
					Err(err).
					Str("contract_title", p.Title).
					Msg("[SEEDER][CONTRACT-DOCUMENT] Failed to insert contract documents")

				errorCount++
				continue
			}
		}

		successCount++
	}

	log.Info().
		Int("success", successCount).
		Int("errors", errorCount).
		Msg("[SEEDER][CONTRACT-DOCUMENT] Seeding completed")

}
