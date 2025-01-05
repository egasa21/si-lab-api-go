package response

import (
	"encoding/json"
	"net/http"
	"github.com/egasa21/si-lab-api-go/internal/errs"
)

// Meta structure that contains meta information.
type Meta struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code,omitempty"`
}

// Response structure that wraps the `meta` and `data` fields.
type Response struct {
	Meta       Meta        `json:"meta"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination,omitempty"`
}

// Pagination structure for paginated responses.
type Pagination struct {
	Page        int `json:"page"`
	PerPage     int `json:"per_page"`
	TotalPages  int `json:"total_pages"`
	TotalItems  int `json:"total_items"`
}

// NewSuccessResponse formats a success response with data.
func NewSuccessResponse(w http.ResponseWriter, data interface{}, message string) {
	response := Response{
		Meta: Meta{
			Success: true,
			Message: message,
		},
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	json.NewEncoder(w).Encode(response)
}

// NewPaginatedSuccessResponse formats a paginated success response.
func NewPaginatedSuccessResponse(w http.ResponseWriter, data interface{}, pagination Pagination, message string) {
	response := Response{
		Meta: Meta{
			Success: true,
			Message: message,
		},
		Data:       data,
		Pagination: pagination,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	json.NewEncoder(w).Encode(response)
}

// NewErrorResponse formats an error response based on AppError.
func NewErrorResponse(w http.ResponseWriter, err *errs.AppError) {
	response := Response{
		Meta: Meta{
			Success:   false,
			Message:   err.Message,
			ErrorCode: err.StatusCode,
		},
		Data: nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode) // HTTP status code
	json.NewEncoder(w).Encode(response)
}
