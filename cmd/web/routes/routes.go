package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stazoloto/rest-api-server/cmd/web/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.IndexHandler())
	app.Get("/snippet", handlers.ShowSnippetHandler())
	app.Get("/snippet/create", handlers.CreateSnippetHandler())
}
