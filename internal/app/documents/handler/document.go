package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *HDocument) Extract(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var err error

	document, err := c.FormFile("document")
	if err != nil {
		return errors.New("please provide the administration document")
	}

	p, err := h.ud.Extract(ctx, document)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(p)
}
