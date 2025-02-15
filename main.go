package main

import (
	"fmt"
	"insan-service-account-backend/database"
	"insan-service-account-backend/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	// Initialize Zap Logger
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Error initializing zap logger: %v\n", err)
	}
	defer logger.Sync() // Flushing logs


	// Start a new Fiber App
	app := fiber.New()
	
	// Custom middleware logger
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()

		// Proceed to the next middleware/handler
		err := c.Next()

		// Log the request details
		duration := time.Since(start)
		status := c.Response().StatusCode()
		method := c.Method()
		path := c.Path()

		// Log based on status code
		switch {
		case status >= 500:
			logger.Error("Server error",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("status", status),
				zap.Duration("duration", duration),
				zap.Error(err),
			)
		case status >= 400:
			logger.Warn("Client error",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("status", status),
				zap.Duration("duration", duration),
			)
		default:
			logger.Info("Request processed",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("status", status),
				zap.Duration("duration", duration),
			)
		}

		return err
	})

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
