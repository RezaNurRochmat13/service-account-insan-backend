package routes

import (
	accountRoutes "insan-service-account-backend/module/account/routes"
	userRoutes "insan-service-account-backend/module/user/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	userRoutes.SetupUserRoutes(api)
	accountRoutes.SetupAccountRoutes(api)
}
