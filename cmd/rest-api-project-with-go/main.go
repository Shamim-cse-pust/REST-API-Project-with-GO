package main

import (
	"fmt"
	"log"

	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/config"
	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
	})

	// Setup all routes
	routes.SetupRoutes(app, cfg)

	// Start server
	fmt.Printf("🚀 Server starting on %s\n", cfg.GetServerAddress())
	log.Fatal(app.Listen(cfg.GetServerAddress()))
}
