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
	api.Get("/book/:id", handlers.GetBook)
	api.Post("/book", handlers.NewBook)
	api.Put("/book:id", handlers.UpdateBook)
	api.Delete("/book:id", handlers.DeleteBook)
}
