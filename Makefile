# Laravel-style migration commands for Go
DB_URL := "mysql://root:123@tcp(localhost:3306)/rest_api_db"
MIGRATIONS_PATH := migrations

.PHONY: migrate rollback rollback-step rollback-all migrate-status migrate-create help

# Run all pending migrations (like php artisan migrate)
migrate:
	@echo "🔄 Running migrations..."
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_URL) up
	@echo "🎉 Migrations completed!"

# Rollback last batch (like php artisan migrate:rollback)
rollback:
	@echo "🔄 Rolling back last batch of migrations..."
	@echo "How many migrations should be rolled back?"
	@echo "1) Rollback 1 migration (latest only)"
	@echo "2) Rollback 3 migrations (if you migrated 3 files in last batch)"
	@echo "3) Custom number"
	@read -p "Choose option [1-3]: " choice; \
	case $$choice in \
		1) migrate -path $(MIGRATIONS_PATH) -database $(DB_URL) down 1 ;; \
		2) migrate -path $(MIGRATIONS_PATH) -database $(DB_URL) down 3 ;; \
		3) read -p "Enter number of migrations to rollback: " steps; \
		   migrate -path $(MIGRATIONS_PATH) -database $(DB_URL) down $$steps ;; \
		*) echo "Invalid choice. Rolling back 1 migration."; \
		   migrate -path $(MIGRATIONS_PATH) -database $(DB_URL) down 1 ;; \
	esac
	@echo "🎉 Rollback completed!"

# Rollback specific number of migrations
rollback-step:
	@echo "🔄 Rolling back specific number of migrations..."
	@read -p "Enter number of steps to rollback: " steps; \
	migrate -path $(MIGRATIONS_PATH) -database $(DB_URL) down $$steps
	@echo "🎉 Rollback completed!"

# Rollback ALL migrations (like php artisan migrate:reset)
rollback-all:
	@echo "🔄 Rolling back ALL migrations..."
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_URL) down
	@echo "🎉 All migrations rolled back!"

# Show migration status (like php artisan migrate:status)
migrate-status:
	@echo "📋 Migration Status:"
	@echo "Available migrations:"
	@ls -1 $(MIGRATIONS_PATH)/*.up.sql | sed 's/.*\///; s/\.up\.sql//'
	@echo ""
	@echo "Current database version:"
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_URL) version || echo "❌ No migrations applied yet"

# Create new migration (like php artisan make:migration)
migrate-create:
	@echo "Creating new migration..."
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) $$name
	@echo "✅ Migration files created!"

# Force migration version (advanced)
migrate-force:
	@read -p "Enter version to force: " version; \
	migrate -path $(MIGRATIONS_PATH) -database $(DB_URL) force $$version

# Show help
help:
	@echo "🚀 Laravel-style Migration Commands:"
	@echo ""
	@echo "make migrate        - Run all pending migrations (like php artisan migrate)"
	@echo "make rollback       - Rollback last batch (like php artisan migrate:rollback)" 
	@echo "make rollback-step  - Rollback specific number of migrations"
	@echo "make rollback-all   - Rollback ALL migrations (like php artisan migrate:reset)"
	@echo "make migrate-status - Show current migration version"
	@echo "make migrate-create - Create new migration files"
	@echo "make migrate-force  - Force set migration version"
	@echo ""
	@echo "Examples:"
	@echo "make migrate"
	@echo "make rollback       # Asks how many migrations to rollback"
	@echo "make rollback-step  # Rollback specific number"
	@echo "make rollback-all   # Rolls back everything"

# Default target
.DEFAULT_GOAL := help
