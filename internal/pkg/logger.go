package pkg

import (
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type LoggerConfig struct {
	GlobalLevel   string
	MinimumLevel  string
	ErrorStack    bool
}

func NewLoggerConfig(globalLevel, minLevel string, errorStack bool) LoggerConfig {
	return LoggerConfig{
		GlobalLevel:   globalLevel,
		MinimumLevel:  minLevel,
		ErrorStack:    errorStack,
	}
}

func NewLogger(config LoggerConfig) zerolog.Logger {
	// Set error stack marshaller if enabled
	if config.ErrorStack {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	}

	// Parse global log level
	level, err := zerolog.ParseLevel(config.GlobalLevel)
	if err != nil {
		level = zerolog.DebugLevel // Default level
	}

	// Set global log level
	zerolog.SetGlobalLevel(level)

	// Parse minimum log level
	minLevel, err := zerolog.ParseLevel(config.MinimumLevel)
	if err != nil {
		minLevel = zerolog.DebugLevel
	}

	// Create logger with minimum level
	logger := zerolog.New(os.Stdout).
		Level(minLevel).
		With().
		Timestamp().
		Logger()

	return logger
}

func ParseBoolFromEnv(envVar string, defaultValue bool) bool {
	val, err := strconv.ParseBool(os.Getenv(envVar))
	if err != nil {
		return defaultValue
	}
	return val
}
