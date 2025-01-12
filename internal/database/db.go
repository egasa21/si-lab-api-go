package database

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/egasa21/si-lab-api-go/configs"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// ConnectDB connects to the database
func ConnectDB(cfg *configs.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool (example)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	// Connection test
	ctx := context.Background()
	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("database connection test failed: %w", err)
	}

	if err := RunMigrations(ctx, cfg); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	fmt.Println("Database connection established")
	return db, nil
}

// RunMigrations runs the database migrations
func RunMigrations(ctx context.Context, cfg *configs.Config) error {
	// Construct the migrations path (example using project root)
	// projectRoot := os.Getenv("PROJECT_ROOT")
	migrationsPath, err := filepath.Abs("../../internal/database/migrations")
	if err != nil {
		return fmt.Errorf("error getting absolute path for migrations: %w", err)
	}

	// Construct the file source URL
	fileSourceURL := "file://" + migrationsPath

	// Construct the database URL
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Initialize the migration
	m, err := migrate.New(fileSourceURL, dbURL)
	if err != nil {
		return fmt.Errorf("failed to initialize migrations: %w", err)
	}

	// Run the migration
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	fmt.Println("Migrations applied successfully")
	return nil
}
