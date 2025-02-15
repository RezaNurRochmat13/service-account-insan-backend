package main

import (
	"insan-service-account-backend/database"
	"insan-service-account-backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new Fiber App
	app := fiber.New()

	// Connect to the database
	database.ConnectDatabase()

	// Setup Routes
	routes.SetupRoutes(app)


	// Send string back for GET calls to the endpoint '/'
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API is up and running")
	})

	// Listen on port 3000
	app.Listen(":3000")
}
