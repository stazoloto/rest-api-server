package static

import "github.com/gofiber/fiber/v2"

func SetupStatic(app *fiber.App) {
	app.Static("/", "rest-api-server/ui")
}
