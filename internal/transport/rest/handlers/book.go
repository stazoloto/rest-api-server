package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	if book.Id == uuid.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No book present", "data": nil})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Book found", "data": book})
}

// NewBook add a new book
func NewBook(ctx *fiber.Ctx) error {
	db := database.DB
	book := new(model.Book)

	err := ctx.BodyParser(book)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	book.Id = uuid.New()
	err = db.Create(&book).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not create book", "data": err})
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "Created Book", "data": book})
}

// UpdateBook update book
func UpdateBook(ctx *fiber.Ctx) error {
	type updateBook struct {
		Title  string `json:"title"`
		Author string `json:"sub_title"`
		Rating string `json:"text"`
	}

	db := database.DB
	var book model.Book

	id := ctx.Params("bookId")
	db.Find(&book, "id = ?", id)

	if book.Id == uuid.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No book present", "data": nil})
	}

	var updateBookData updateBook
	err := ctx.BodyParser(&updateBookData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	book.Title = updateBookData.Title
	book.Author = updateBookData.Author
	book.Rating = updateBookData.Rating

	db.Save(&book)

	return ctx.JSON(fiber.Map{"status": "success", "message": "Book Updated", "data": book})
}

// DeleteBook delete book
func DeleteBook(ctx *fiber.Ctx) error {
	db := database.DB
	var book model.Book

	id := ctx.Params("bookId")
	db.Find(&book, "id = ?", id)

	if book.Id == uuid.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No book present", "data": nil})
	}

	err := db.Delete(&book, "id = ?", id).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Failed to delete book", "data": nil})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Book Deleted"})
}
