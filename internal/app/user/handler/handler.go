package handler

import (
	"time"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/contract"
	"github.com/Clausia-Ifest/clausia-server/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

type HUser struct {
	v  *validator.Validate
	m  middleware.IMiddleware
	uu contract.IUUser
}

func New(v *validator.Validate, m middleware.IMiddleware, uu contract.IUUser) *HUser {
	return &HUser{
		v:  v,
		m:  m,
		uu: uu,
	}
}

func (h *HUser) MountRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("", timeout.NewWithContext(h.Auth, 3*time.Second))

	auth.Get("", h.m.RequiredAuth, timeout.NewWithContext(h.Self, 3*time.Second))
}
