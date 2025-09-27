package handler

import (
	"time"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/contract"
	"github.com/Clausia-Ifest/clausia-server/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

type HDocument struct {
	v  *validator.Validate
	m  middleware.IMiddleware
	ud contract.IUDocument
}

func New(v *validator.Validate, m middleware.IMiddleware, ud contract.IUDocument) *HDocument {
	return &HDocument{
		v:  v,
		m:  m,
		ud: ud,
	}
}

func (h *HDocument) MountRoutes(router fiber.Router) {
	documents := router.Group("/documents")

	documents.Use(h.m.RequiredAuth)
	documents.Post("/", h.m.RequiredRoles("Admin"), timeout.NewWithContext(h.Extract, 30*time.Second))
}
