package transactionController

import (
	transactionModel "insan-service-account-backend/module/transaction/model"
	transactionService "insan-service-account-backend/module/transaction/service"

	"github.com/gofiber/fiber/v2"
)

func GetAllTransactions(c *fiber.Ctx) error {
	transactions, err := transactionService.FindAllTransactions()

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to find transactions", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Transactions Found", "data": transactions})
}

func GetTransactionByAccountID(c *fiber.Ctx) error {
	transaction, err := transactionService.FindTransactionByAccountNumber(c.Params("accountNumber"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to find transactions", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Transactions Found", "data": transaction})
}

func CreateNewTransaction(c *fiber.Ctx) error {
	transaction := new(transactionModel.Transaction)

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to parse request body", "data": err.Error()})
	}

	transaction, err := transactionService.CreateNewTransaction(transaction)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to create transaction", "data": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Transaction Created", "data": transaction})
}
