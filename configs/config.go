package configs

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	AppPort       string
	LogLevel      string
	LogErrorStack bool
}

func LoadConfig() *Config {

	if os.Getenv("DYNO") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	// Parse boolean for LOG_ERROR_STACK
	logErrorStack := false
	if strings.ToLower(os.Getenv("LOG_ERROR_STACK")) == "true" {
		logErrorStack = true
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		AppPort:       port,
		LogLevel:      os.Getenv("LOG_LEVEL"),
		LogErrorStack: logErrorStack,
	}
}
