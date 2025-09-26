package middleware

import (
	"github.com/Clausia-Ifest/clausia-server/pkg/token"
	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	token token.IJWT
}

type IMiddleware interface {
	RequiredAuth(c *fiber.Ctx) error
	RequiredRoles(roles ...string) fiber.Handler
}

func New(token token.IJWT) IMiddleware {
	return &Middleware{
		token: token,
	}
}
