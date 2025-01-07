package pkg

import "net/http"

type AppError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewAppError(message string, statusCode int) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (e *AppError) Error() string {
	return e.Message
}


var (
	ErrBadRequest       = NewAppError("Bad request", http.StatusBadRequest)
	ErrNotFound         = NewAppError("Resource not found", http.StatusNotFound)
	ErrInternalServer   = NewAppError("Internal server error", http.StatusInternalServerError)
	ErrUnauthorized     = NewAppError("Unauthorized", http.StatusUnauthorized)
	ErrForbidden        = NewAppError("Forbidden", http.StatusForbidden)
)
