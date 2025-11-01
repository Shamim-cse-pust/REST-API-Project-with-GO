package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/config"
	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/routes"
	"github.com/gofiber/fiber/v2"
)

// setupFileLogging configures logging to write to logs/app.log file
func setupFileLogging() {
	// Create logs directory if it doesn't exist
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Printf("❌ Failed to create logs directory: %v", err)
		return
	}

	// Create app.log file in logs folder
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("❌ Failed to create logs/app.log file: %v", err)
		return
	}

	// Set log output to both console and file
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	// Set log format with timestamp
	log.SetFlags(log.LstdFlags)
}

func main() {
	// Setup file logging
	setupFileLogging()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
	})

	// Setup all routes
	routes.SetupRoutes(app, cfg)

	// Start server in background goroutine
	go func() {
		log.Printf("🚀 %s v%s starting on %s", cfg.App.Name, cfg.App.Version, cfg.GetServerAddress())
		if err := app.Listen(cfg.GetServerAddress()); err != nil {
			log.Printf("❌ Server error: %v", err)
		}
	}()

	// Create channel to listen for interrupt signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until we receive a signal
	<-c

	// Graceful shutdown with context
	log.Println("🛑 Shutting down gracefully...")

	// Give active requests 10 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("❌ Server shutdown error: %v", err)
	}
	log.Println("✅ Server stopped")
}
