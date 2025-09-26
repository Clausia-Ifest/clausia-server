package handler

import (
	"errors"
	"net/http"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *HContract) SubmitContract(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var req dto.SubmitContractRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	var err error

	req.AdministrationDocument, err = c.FormFile("administration_document")
	if err != nil {
		return errors.New("administration document not found")
	}

	req.LegalDocument, err = c.FormFile("legal_document")
	if err != nil {
		return errors.New("legal document not found")
	}

	req.TechnicalDocument, err = c.FormFile("technical_document")
	if err != nil {
		return errors.New("technical document not found")
	}

	req.FinancialDocument, err = c.FormFile("financial_document")
	if err != nil {
		return errors.New("financial document not found")
	}

	if err := h.v.Struct(req); err != nil {
		return err
	}

	if err := h.uc.Submit(ctx, req); err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "OK",
	})
}

func (h *HContract) All(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var req dto.AllContractsRequest

	req.Limit = int64(c.QueryInt("limit", 10))
	req.Page = int64(c.QueryInt("page", 1))

	p, err := h.uc.All(ctx, req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(p)
}

func (h *HContract) Update(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var req dto.UpdateContractRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	var err error
	req.ID, err = uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	if err := h.uc.Update(ctx, req); err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "OK",
	})
}
