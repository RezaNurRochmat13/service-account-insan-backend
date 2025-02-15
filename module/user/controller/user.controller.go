package userController

import (
	"insan-service-account-backend/database"
	userModel "insan-service-account-backend/module/user/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []userModel.User

	db.Find(&users)

	return c.JSON(fiber.Map{"status": "success", "message": "Users Found", "data": users})
}

func CreateNewUser(c *fiber.Ctx) error {
	db := database.DB
    user := new(userModel.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to parse request body", "data": err})
	}

	var existingUser userModel.User
	if err := db.Where("phone_number = ?", user.PhoneNumber).First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "User with this phone number already exists", "data": nil})
	}

	user.ID = uuid.New()

	if err := db.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User Created", "data": user})
}
