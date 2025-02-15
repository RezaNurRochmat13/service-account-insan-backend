package accountRoutes

import (
	accountController "insan-service-account-backend/module/account/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupAccountRoutes(router fiber.Router) {
	account := router.Group("/accounts")
	
	// Get current balance
	account.Get("/:id", accountController.GetAccountCurrentBalance)
	// Create a Account
	account.Post("/", accountController.CreateNewAccount)
}
