package main

import (
	"os"
	"strings"

	"github.com/egasa21/si-lab-api-go/configs"
	"github.com/egasa21/si-lab-api-go/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Load environment variables
	cfg := configs.LoadConfig()

	// Set up zerolog
	logLevel := strings.ToLower(cfg.LogLevel)
	zerolog.SetGlobalLevel(parseLogLevel(logLevel))

	// Configure logger to write to stdout
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("service", "si-lab-api-go").
		Logger()

	// Initialize the server with logger
	srv := server.NewServer(cfg, logger)

	// Start the server
	if err := srv.Start(); err != nil {
		log.Fatal().Err(err).Msg("Error starting server")
	}
}

// Helper function to parse log level
func parseLogLevel(level string) zerolog.Level {
	switch level {
	case "panic":
		return zerolog.PanicLevel
	case "fatal":
		return zerolog.FatalLevel
	case "error":
		return zerolog.ErrorLevel
	case "warn":
		return zerolog.WarnLevel
	case "info":
		return zerolog.InfoLevel
	case "debug":
		return zerolog.DebugLevel
	case "trace":
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}
