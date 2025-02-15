package routes

import (
	accountRoutes "insan-service-account-backend/module/account/routes"
	transactionRoutes "insan-service-account-backend/module/transaction/routes"
	userRoutes "insan-service-account-backend/module/user/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiGroup := app.Group("/api/v1")

	// Register routes
	userRoutes.SetupUserRoutes(apiGroup)
	accountRoutes.SetupAccountRoutes(apiGroup)
	transactionRoutes.SetupTransactionRoutes(apiGroup)
}
