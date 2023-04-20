package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func CreateSnippetHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("create snippet")
	}
}
