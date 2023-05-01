package handlers

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/stazoloto/rest-api-server/internal/database/postgresql"
	"github.com/stazoloto/rest-api-server/internal/model"
)

// GetBooks show all books
func GetBooks(ctx *fiber.Ctx) error {
	db := database.DB
	var books []model.Book

	db.Find(&books)
	if len(books) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No books present", "data": nil})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Books found", "data": books})
}

// GetBook show one book
func GetBook(ctx *fiber.Ctx) error {
	db := database.DB
	var book model.Book

	id := ctx.Params("bookId")
	db.Find(&book, "id = ?", id)
}

// NewBook add a new book
func NewBook(ctx *fiber.Ctx) error {
	return ctx.SendString("Add a new book")
}

// UpdateBook update book
func UpdateBook(ctx *fiber.Ctx) error {
	return ctx.SendString("Update book")
}

// DeleteBook delete book
func DeleteBook(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete book")
}
