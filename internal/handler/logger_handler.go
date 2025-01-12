package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type LoggerState struct {
	LoggerMinimumLevel string `json:"logger_minimum_level"`
	GlobalLogLevel     string `json:"global_log_level"`
	LogErrorStack      bool   `json:"log_error_stack"`
}

var loggerState = &LoggerState{
	LoggerMinimumLevel: "debug",
	GlobalLogLevel:     "debug",
	LogErrorStack:      false,
}

var loggerStateMutex sync.RWMutex

func GetLoggerState(w http.ResponseWriter, r *http.Request) {
	loggerStateMutex.RLock()
	defer loggerStateMutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loggerState)
}

func UpdateLoggerState(w http.ResponseWriter, r *http.Request) {
	var newState LoggerState
	if err := json.NewDecoder(r.Body).Decode(&newState); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	loggerStateMutex.Lock()
	defer loggerStateMutex.Unlock()

	loggerState.LoggerMinimumLevel = newState.LoggerMinimumLevel
	loggerState.GlobalLogLevel = newState.GlobalLogLevel
	loggerState.LogErrorStack = newState.LogErrorStack

	// Apply the changes dynamically
	zerolog.SetGlobalLevel(parseLogLevel(newState.GlobalLogLevel))
	zerolog.ErrorStackMarshaler = nil
	if newState.LogErrorStack {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loggerState)
}

func parseLogLevel(level string) zerolog.Level {
	parsed, err := zerolog.ParseLevel(level)
	if err != nil {
		return zerolog.DebugLevel
	}
	return parsed
}
