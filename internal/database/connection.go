package database

import (
	"log"

	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// database holds the GORM database connection
var database *gorm.DB

// ConnectDatabase creates a new GORM database connection
func ConnectDatabase(cfg *config.Config) error {
	// Create connection string for GORM
	dsn := cfg.GetDatabaseURL()

	// Configure GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable SQL logging
	})

	if err != nil {
		return err
	}

	// Get underlying sql.DB for connection pool configuration
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return err
	}

	// Set global database variable
	database = db

	log.Println("✅ Database connected successfully with GORM")

	return nil
}

// GetDB returns the GORM database instance
func GetDB() *gorm.DB {
	return database
}

// CloseDatabase closes the database connection
func CloseDatabase() {
	if database != nil {
		if sqlDB, err := database.DB(); err == nil {
			log.Println("🔐 Closing database connection...")
			sqlDB.Close()
		}
	}
}
