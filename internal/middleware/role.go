package middleware

import (
	"errors"
	"slices"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) RequiredRoles(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("user.role").(string)
		if !ok {
			return errors.New("failed to get user's role")
		}

		if !slices.Contains(roles, role) {
			return errors.New("ur not allowed to access this endpoint")
		}

		return c.Next()
	}
}
