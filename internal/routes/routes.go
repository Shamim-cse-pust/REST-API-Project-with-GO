package routes

import (
	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/config"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, cfg *config.Config) {
	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":     "Welcome to REST API with Go & Fiber!",
			"app":         cfg.App.Name,
			"version":     cfg.App.Version,
			"environment": cfg.App.Environment,
		})
	})

	// Health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "healthy",
			"database":  "connected", // TODO: Add real DB health check
			"timestamp": c.Context().Time(),
		})
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World",
		})
	})
}
