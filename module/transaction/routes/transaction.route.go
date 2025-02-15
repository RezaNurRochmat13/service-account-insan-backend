package transactionRoutes

import (
	transactionController "insan-service-account-backend/module/transaction/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupTransactionRoutes(router fiber.Router) {
	transaction := router.Group("/transactions")
	
	transaction.Get("/", transactionController.GetAllTransactions)
	transaction.Get("/account/:accountNumber", transactionController.GetTransactionByAccountID)
	transaction.Post("/", transactionController.CreateNewTransaction)
}
