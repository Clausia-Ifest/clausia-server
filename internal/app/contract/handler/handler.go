package handler

import (
	"time"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/contract"
	"github.com/Clausia-Ifest/clausia-server/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

type HContract struct {
	v  *validator.Validate
	m  middleware.IMiddleware
	uc contract.IUContract
}

func New(v *validator.Validate, m middleware.IMiddleware, uc contract.IUContract) *HContract {
	return &HContract{
		v:  v,
		m:  m,
		uc: uc,
	}
}

func (h *HContract) MountRoutes(router fiber.Router) {
	contracts := router.Group("/contracts")

	contracts.Use(h.m.RequiredAuth)
	contracts.Post("/", h.m.RequiredRoles("Admin"), timeout.NewWithContext(h.SubmitContract, 10*time.Second))
	contracts.Get("/", h.m.RequiredRoles("Admin", "Legal", "Manager"), timeout.NewWithContext(h.All, 3*time.Second))
	contracts.Get("/:id", h.m.RequiredRoles("Legal"), timeout.NewWithContext(h.Get, 3*time.Second))
	contracts.Patch("/:id", h.m.RequiredRoles("Legal", "Manager"), timeout.NewWithContext(h.Update, 3*time.Second))

	logs := contracts.Group("/logs")
	logs.Get("/", h.m.RequiredRoles("Admin"), timeout.NewWithContext(h.Logs, 3*time.Second))

	contracts.Post("/:id/chat", h.m.RequiredRoles("Legal"), timeout.NewWithContext(h.Chat, 3*time.Second))
}
