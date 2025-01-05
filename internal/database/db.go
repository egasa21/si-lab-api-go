package database

import (
	"database/sql"
	"fmt"
	"log"

	// "os"
	// "path/filepath"

	"github.com/egasa21/si-lab-api-go/configs"
	// "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// ConnectDB connects to the database
func ConnectDB(cfg *configs.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// connection test
	if err = db.Ping(); err != nil {
		log.Fatalf("Database connection test failed: %v", err)
	}

	fmt.Println("Database connection established")
	return db
}

// RunMigrations runs the database migrations
// func RunMigrations(cfg *configs.Config) {
// 	// Get the working directory
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		log.Fatalf("Error getting working directory: %v", err)
// 	}

// 	// Construct the migrations path
// 	relativePath := filepath.Join(cwd, "../../internal/database/migrations")
// 	migrationsPath, err := filepath.Abs(relativePath)
// 	if err != nil {
// 		log.Fatalf("Error getting absolute path for migrations: %v", err)
// 	}

// 	// Construct the file source URL
// 	fileSourceURL := "file://" + migrationsPath

// 	// Construct the database URL
// 	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
// 		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

// 	// Initialize the migration
// 	m, err := migrate.New(fileSourceURL, dbURL)
// 	if err != nil {
// 		log.Fatalf("Failed to initialize migrations: %v", err)
// 	}

// 	// Run the migration
// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatalf("Failed to run migrations: %v", err)
// 	}

// 	fmt.Println("Migrations applied successfully")
// }
