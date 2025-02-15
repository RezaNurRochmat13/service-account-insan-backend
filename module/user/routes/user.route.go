package userRoutes

import (
	userController "insan-service-account-backend/module/user/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	account := router.Group("/users")
	
	account.Get("/", userController.GetAllUsers)
	account.Post("/", userController.CreateNewUser)
}
