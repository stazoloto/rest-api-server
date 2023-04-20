package handlers

import "github.com/gofiber/fiber/v2"

func ShowSnippetHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Snippet")
	}
}
