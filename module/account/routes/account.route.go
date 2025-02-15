package accountRoutes

import (
	accountController "insan-service-account-backend/module/account/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupAccountRoutes(router fiber.Router) {
	account := router.Group("/accounts")
	
	account.Get("/", accountController.GetAllAccounts)
	account.Get("/balance/:accountNumber", accountController.GetAccountCurrentBalance)
	account.Post("/", accountController.CreateNewAccount)
}
