package routes

import (
	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/config"
	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/database"
	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/handlers"
	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/repositories"
	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/services"
	"github.com/gofiber/fiber/v2"
)

// newUserHandler creates a fully configured user handler with all dependencies
func newUserHandler() *handlers.UserHandler {
	db := database.GetDB()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	return handlers.NewUserHandler(userService)
}

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

	// Initialize handlers using factory
	userHandler := newUserHandler()

	// API routes
	v1 := app.Group("/api/v1")
	users := v1.Group("/users")
	users.Get("/", userHandler.GetAllUsers)
	users.Post("/", userHandler.CreateUser)
	users.Get("/:id", userHandler.GetUserByID)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
}
