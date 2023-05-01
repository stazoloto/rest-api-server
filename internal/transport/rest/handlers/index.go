package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func IndexHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Home")
}
