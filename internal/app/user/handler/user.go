package handler

import (
	"net/http"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *HUser) Auth(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var req dto.SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.v.Struct(req); err != nil {
		return err
	}

	p, err := h.uu.Auth(ctx, req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(p)
}

func (h *HUser) Self(c *fiber.Ctx) error {
	ctx := c.UserContext()

	uuids := "01997fb9-c4d5-754b-a0d6-e622347f66f6"

	uuid, err := uuid.Parse(uuids)
	if err != nil {
		return err
	}

	p, err := h.uu.Self(ctx, uuid)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(p)
}
