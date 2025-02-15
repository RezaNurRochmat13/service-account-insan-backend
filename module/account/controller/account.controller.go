package accountController

import (
	"insan-service-account-backend/database"
	accountModel "insan-service-account-backend/module/account/model"
	"insan-service-account-backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAccountCurrentBalance(c *fiber.Ctx) error {
    db := database.DB
    var account accountModel.Account

    accountNumber := c.Params("accountNumber")

    if err := db.Find(&account, "account_number = ?", accountNumber).Error; err != nil {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Account not found", "data": err})
    }

    return c.JSON(fiber.Map{"status": "success", "message": "Account Found", "data": account})
}

func CreateNewAccount(c *fiber.Ctx) error {
    db := database.DB
    account := new(accountModel.Account)

    if err := c.BodyParser(account); err != nil {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to create account", "data": err})
    }

    account.ID = uuid.New()
    account.AccountNumber = utils.GenerateAccountNumber(10)
    
    if err := db.Create(&account).Error; err != nil {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to create account", "data": err})
    }

    if err := db.Find(&account, "account_number = ?", account.AccountNumber).Error; err == nil {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Account already exists", "data": err})
    }

    return c.JSON(fiber.Map{"status": "success", "message": "Account Created", "data": account})
}
