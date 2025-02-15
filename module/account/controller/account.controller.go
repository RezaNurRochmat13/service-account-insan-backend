package accountController

import (
	accountModel "insan-service-account-backend/module/account/model"
	accountService "insan-service-account-backend/module/account/service"

	"github.com/gofiber/fiber/v2"
)

func GetAllAccounts(c *fiber.Ctx) error {
    accounts, err := accountService.FindAllAccounts()

    if err != nil {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to find accounts", "data": nil})
    }

    return c.JSON(fiber.Map{"status": "success", "message": "Accounts Found", "data": accounts})
}

func GetAccountCurrentBalance(c *fiber.Ctx) error {
    accountNumber := c.Params("accountNumber")
    account, err := accountService.FindAccountByAccountNumber(accountNumber)

    if err != nil {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to find account", "data": nil})
    }

    return c.JSON(fiber.Map{"status": "success", "message": "Account Found", "data": account})
}

func CreateNewAccount(c *fiber.Ctx) error {
    account := new(accountModel.Account)

    if err := c.BodyParser(account); err != nil {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to create account", "data": err})
    }

    account, err := accountService.SaveAccounts(account)

    if err != nil {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to create account", "data": err})
    }

    return c.JSON(fiber.Map{"status": "success", "message": "Account Created", "data": account})
}
