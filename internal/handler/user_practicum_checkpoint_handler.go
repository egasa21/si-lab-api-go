package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/egasa21/si-lab-api-go/internal/dto"
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/pkg"
	"github.com/egasa21/si-lab-api-go/internal/pkg/response"
	"github.com/egasa21/si-lab-api-go/internal/service"
)

type UserPracticumCheckpointHandler struct {
	service service.UserPracticumCheckpointService
}

func NewUserPracticumCheckpointHandler(service service.UserPracticumCheckpointService) *UserPracticumCheckpointHandler {
	return &UserPracticumCheckpointHandler{service: service}
}

// CreateCheckpoint handles the creation of a new user practicum checkpoint
func (h *UserPracticumCheckpointHandler) CreateCheckpoint(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserPracticumCheckpointRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	checkpoint := model.UserPracticumCheckpoint{
		UserID:      req.UserID,
		PracticumID: req.PracticumID,
		ModuleID:    req.ModuleID,
		ContentID:   req.ContentID,
	}

	// Call the service to create the checkpoint
	err = h.service.CreateCheckpoint(&checkpoint)
	if err != nil {
		appErr := pkg.NewAppError("Failed to create checkpoint", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "Checkpoint created successfully")
}

// GetCheckpointByUserAndPracticum retrieves a checkpoint by user ID and practicum ID
func (h *UserPracticumCheckpointHandler) GetCheckpointByUserAndPracticum(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid user ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	practicumID, err := strconv.Atoi(r.URL.Query().Get("practicum_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid practicum ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	// Retrieve the checkpoint from the service
	checkpoint, err := h.service.GetCheckpointByUserAndPracticum(userID, practicumID)
	if err != nil {
		appErr := pkg.NewAppError("Unable to fetch checkpoint", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	if checkpoint == nil {
		appErr := pkg.NewAppError("Checkpoint not found", http.StatusNotFound)
		response.NewErrorResponse(w, appErr)
		return
	}

	// Respond with the checkpoint data
	response.NewSuccessResponse(w, checkpoint, "Checkpoint retrieved successfully")
}

// UpdateCheckpoint handles updating a user practicum checkpoint
func (h *UserPracticumCheckpointHandler) UpdateCheckpoint(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserPracticumCheckpointRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	checkpoint := model.UserPracticumCheckpoint{
		UserID:      req.UserID,
		PracticumID: req.PracticumID,
		ModuleID:    req.ModuleID,
		ContentID:   req.ContentID,
	}

	err = h.service.UpdateCheckpoint(&checkpoint)
	if err != nil {
		appErr := pkg.NewAppError("Failed to update checkpoint", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	// Respond with success
	response.NewSuccessResponse(w, nil, "Checkpoint updated successfully")
}

// DeleteCheckpoint handles the deletion of a checkpoint by ID
func (h *UserPracticumCheckpointHandler) DeleteCheckpoint(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid checkpoint ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	err = h.service.DeleteCheckpoint(id)
	if err != nil {
		appErr := pkg.NewAppError("Failed to delete checkpoint", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "Checkpoint deleted successfully")
}
