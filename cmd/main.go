package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stazoloto/rest-api-server/cmd/routes"
)

func initDatabase() {
	var err error
	database.DBConn := 
}

func main() {
	app := fiber.New()

	router := &routes.Router{}
	router.IndexSetupRoutes(app)
	router.BookSetupRoutes(app)

	app.Listen("localhost:3000")
}
