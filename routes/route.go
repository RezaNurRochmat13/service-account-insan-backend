package routes

import (
	noteRoutes "golang-boilerplate-example/module/note/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	noteRoutes.SetupNoteRoutes(api)
}
