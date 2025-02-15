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

    id := c.Params("id")

    db.Find(&account, "id = ?", id)

    if account.ID == uuid.Nil {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No account present", "data": nil})
    }

    return c.JSON(fiber.Map{"status": "success", "message": "Account Found", "data": account})
}

func CreateNewAccount(c *fiber.Ctx) error {
    db := database.DB
    account := new(accountModel.Account)

    if err := c.BodyParser(account); err != nil {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to parse request body", "data": err})
    }

    account.ID = uuid.New()
    account.AccountNumber = utils.GenerateAccountNumber(10)
    err := db.Create(&account).Error

    if err != nil {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to create account", "data": err})
    }

    return c.JSON(fiber.Map{"status": "success", "message": "Account Created", "data": account})
}
