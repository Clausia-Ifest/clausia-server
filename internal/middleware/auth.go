package middleware

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) RequiredAuth(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	if header == "" {
		return errors.New("no bearer provided")
	}

	headerSlice := strings.Split(header, " ")
	if len(headerSlice) != 2 || headerSlice[0] != "Bearer" {
		return errors.New("invalid bearer token")
	}

	token := headerSlice[1]
	validToken, err := m.token.Decode(token)
	if err != nil {
		return err
	}

	c.Locals("user.id", validToken.User.ID)
	c.Locals("user.role", validToken.User.Role)

	return c.Next()
}
