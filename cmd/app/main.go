package main

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/stazoloto/rest-api-server/internal/database/postgresql"
	"github.com/stazoloto/rest-api-server/internal/transport/rest/routes"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router := &routes.Router{}
	router.IndexSetupRoutes(app)
	router.BookSetupRoutes(app)

	app.Listen("localhost:3000")
}
