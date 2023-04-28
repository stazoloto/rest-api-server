package handlers

import "github.com/gofiber/fiber/v2"

func GetBooks(ctx *fiber.Ctx) error {
	return ctx.SendString("All Books")
}

func GetBook(ctx *fiber.Ctx) error {
	return ctx.SendString("Single Book")
}

func NewBook(ctx *fiber.Ctx) error {
	return ctx.SendString("Add a new book")
}

func DeleteBook(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete book")
}
