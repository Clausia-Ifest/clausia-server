package seeder

import (
	"time"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/entity"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/enum"
	"github.com/google/uuid"
)

var Contracts = []entity.Contract{
	{
		ID:                uuid.New(),
		HumanID:           "CTR-2025-OK",
		Email:             "john.doe@example.com",
		Title:             "Service Agreement",
		Company:           "PT. Alpha Tech",
		RiskLevel:         enum.RiskLevelLow,
		Status:            enum.StatusWaiting,
		ApplicationStatus: enum.ASManager,
		Category:          enum.CategoryLicenseSoftwareAgreement,
		StartDate:         time.Now().AddDate(0, -1, 0),
		EndDate:           time.Now().AddDate(1, 0, 0),

		ContractDocument: []entity.ContractDocument{
			{
				DocumentHash: "c00e39714fbf9aefcefd8216a03ae29028aed996c2ed8a5606f3238c728123cf",
				URL:          "",
				Category:     enum.DocumentCategoryAdministration,
			},
			{
				DocumentHash: "69616d6763c89c777368b7d396d3383e9da925d5e54fc962616e6e8584176729",
				URL:          "",
				Category:     enum.DocumentCategoryLegal,
			},
			{
				DocumentHash: "31367d11d05e5bb8ec59d310e777064710b980fa35b577028239f09fa2db0c66",
				URL:          "",
				Category:     enum.DocumentCategoryTechnical,
			},
			{
				DocumentHash: "ec9c9d95fdf188012c325b908dab7d0ee771e2fcc22761fdb66fcfe042a93840",
				URL:          "",
				Category:     enum.DocumentCategoryFinancial,
			},
		},
	},
	{
		ID:                uuid.New(),
		HumanID:           "CTR-2025-OK",
		Email:             "john.doe@example.com",
		Title:             "Service Agreement",
		Company:           "PT. Alpha Tech",
		RiskLevel:         enum.RiskLevelLow,
		Status:            enum.StatusWaiting,
		ApplicationStatus: enum.ASManager,
		Category:          enum.CategoryLicenseSoftwareAgreement,
		StartDate:         time.Now().AddDate(0, -1, 0),
		EndDate:           time.Now().AddDate(1, 0, 0),

		ContractDocument: []entity.ContractDocument{
			{
				DocumentHash: "c00e39714fbf9aefcefd8216a03ae29028aed996c2ed8a5606f3238c728123cf",
				URL:          "",
				Category:     enum.DocumentCategoryAdministration,
			},
			{
				DocumentHash: "69616d6763c89c777368b7d396d3383e9da925d5e54fc962616e6e8584176729",
				URL:          "",
				Category:     enum.DocumentCategoryLegal,
			},
			{
				DocumentHash: "31367d11d05e5bb8ec59d310e777064710b980fa35b577028239f09fa2db0c66",
				URL:          "",
				Category:     enum.DocumentCategoryTechnical,
			},
			{
				DocumentHash: "ec9c9d95fdf188012c325b908dab7d0ee771e2fcc22761fdb66fcfe042a93840",
				URL:          "",
				Category:     enum.DocumentCategoryFinancial,
			},
		},
	},
	{
		ID:                uuid.New(),
		HumanID:           "CTR-2025-OK",
		Email:             "john.doe@example.com",
		Title:             "Service Agreement",
		Company:           "PT. Alpha Tech",
		RiskLevel:         enum.RiskLevelLow,
		Status:            enum.StatusWaiting,
		ApplicationStatus: enum.ASManager,
		Category:          enum.CategoryLicenseSoftwareAgreement,
		StartDate:         time.Now().AddDate(0, -1, 0),
		EndDate:           time.Now().AddDate(1, 0, 0),

		ContractDocument: []entity.ContractDocument{
			{
				DocumentHash: "c00e39714fbf9aefcefd8216a03ae29028aed996c2ed8a5606f3238c728123cf",
				URL:          "",
				Category:     enum.DocumentCategoryAdministration,
			},
			{
				DocumentHash: "69616d6763c89c777368b7d396d3383e9da925d5e54fc962616e6e8584176729",
				URL:          "",
				Category:     enum.DocumentCategoryLegal,
			},
			{
				DocumentHash: "31367d11d05e5bb8ec59d310e777064710b980fa35b577028239f09fa2db0c66",
				URL:          "",
				Category:     enum.DocumentCategoryTechnical,
			},
			{
				DocumentHash: "ec9c9d95fdf188012c325b908dab7d0ee771e2fcc22761fdb66fcfe042a93840",
				URL:          "",
				Category:     enum.DocumentCategoryFinancial,
			},
		},
	},
	{
		ID:                uuid.New(),
		HumanID:           "CTR-2025-OK",
		Email:             "john.doe@example.com",
		Title:             "Service Agreement",
		Company:           "PT. Alpha Tech",
		RiskLevel:         enum.RiskLevelLow,
		Status:            enum.StatusWaiting,
		ApplicationStatus: enum.ASManager,
		Category:          enum.CategoryLicenseSoftwareAgreement,
		StartDate:         time.Now().AddDate(0, -1, 0),
		EndDate:           time.Now().AddDate(1, 0, 0),

		ContractDocument: []entity.ContractDocument{
			{
				DocumentHash: "c00e39714fbf9aefcefd8216a03ae29028aed996c2ed8a5606f3238c728123cf",
				URL:          "",
				Category:     enum.DocumentCategoryAdministration,
			},
			{
				DocumentHash: "69616d6763c89c777368b7d396d3383e9da925d5e54fc962616e6e8584176729",
				URL:          "",
				Category:     enum.DocumentCategoryLegal,
			},
			{
				DocumentHash: "31367d11d05e5bb8ec59d310e777064710b980fa35b577028239f09fa2db0c66",
				URL:          "",
				Category:     enum.DocumentCategoryTechnical,
			},
			{
				DocumentHash: "ec9c9d95fdf188012c325b908dab7d0ee771e2fcc22761fdb66fcfe042a93840",
				URL:          "",
				Category:     enum.DocumentCategoryFinancial,
			},
		},
	},
}
