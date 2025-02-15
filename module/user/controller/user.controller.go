package userController

import (
	userModel "insan-service-account-backend/module/user/model"
	userService "insan-service-account-backend/module/user/service"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	users, err := userService.FindAllUsers()

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to find users", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Users Found", "data": users})
}

func CreateNewUser(c *fiber.Ctx) error {
    user := new(userModel.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to parse request body", "data": err})
	}

	user, err := userService.SaveUser(user)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User Created", "data": user})
}
