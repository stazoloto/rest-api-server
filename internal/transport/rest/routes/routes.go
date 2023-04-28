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
	app.Get("/api/v1/book", handlers.GetBooks)
	app.Get("/api/v1/book/:id", handlers.GetBook)
	app.Post("/api/v1/book", handlers.NewBook)
	app.Delete("/api/v1/book:id", handlers.DeleteBook)
}
