package handler

import (
	"encoding/json"
	"net/http"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/pkg"
	"github.com/egasa21/si-lab-api-go/internal/pkg/response"
	"github.com/egasa21/si-lab-api-go/internal/service"
)

type PracticumHandler struct {
	service service.PracticumService
}

func NewPracticumHandler(service service.PracticumService) *PracticumHandler {
	return &PracticumHandler{service: service}
}

func (h *PracticumHandler) CreatePracticum(w http.ResponseWriter, r *http.Request) {
	var practicum model.Practicum
	err := json.NewDecoder(r.Body).Decode(&practicum)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	// check existing practicum

	err = h.service.CreatePracticum(&practicum)
	if err != nil {
		appErr := pkg.NewAppError("Failed to create practicum", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "Practicum created successfully")
}
