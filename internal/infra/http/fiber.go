package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func NewFiber(appEnv string) *fiber.App {
	enablePrint := true
	if appEnv == "prod" || appEnv == "production" {
		enablePrint = false
	}

	return fiber.New(fiber.Config{
		EnablePrintRoutes: enablePrint,
		ErrorHandler:      errorHandler,
		BodyLimit:         100 * 1024 * 1024,
	})
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
