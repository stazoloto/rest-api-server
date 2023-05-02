package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stazoloto/rest-api-server/internal/transport/rest/handlers"
)

type Router struct{}

func (r *Router) IndexSetupRoutes(app *fiber.App) {
	app.Get("/", handlers.IndexHandler)
}

func (r *Router) BookSetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/book", handlers.GetBooks)
	api.Get("/book/:bookId", handlers.GetBook)
	api.Post("/book", handlers.NewBook)
	api.Put("/book/:bookId", handlers.UpdateBook)
	api.Delete("/book/:bookId", handlers.DeleteBook)
}
