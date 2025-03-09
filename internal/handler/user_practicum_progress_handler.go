package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/egasa21/si-lab-api-go/internal/dto"
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/pkg"
	"github.com/egasa21/si-lab-api-go/internal/pkg/response"
	"github.com/egasa21/si-lab-api-go/internal/service"
)

type UserPracticumProgressHandler struct {
	service service.UserPracticumProgressService
}

func NewUserPracticumProgressHandler(service service.UserPracticumProgressService) *UserPracticumProgressHandler {
	return &UserPracticumProgressHandler{service: service}
}

func (h *UserPracticumProgressHandler) CreateProgress(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserPracticumProgressRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	progress := model.UserPracticumProgress{
		UserID:      req.UserID,
		PracticumID: req.PracticumID,
		Progress:    req.Progress,
	}

	err = h.service.CreateProgress(&progress)
	if err != nil {
		appErr := pkg.NewAppError("Failed to create progress", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "User practicum progress created successfully")
}

func (h *UserPracticumProgressHandler) GetProgress(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid user ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	practicumID, err := strconv.Atoi(r.PathValue("practicum_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid practicum ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	progress, err := h.service.GetProgress(userID, practicumID)
	if err != nil {
		appErr := pkg.NewAppError("Progress not found", http.StatusNotFound)
		response.NewErrorResponse(w, appErr)
		return
	}

	var completedAt *string
	if progress.CompletedAt != nil {
		formatted := progress.CompletedAt.Format(time.RFC3339)
		completedAt = &formatted
	}

	res := dto.UserPracticumProgressResponse{
		ID:          progress.ID,
		UserID:      progress.UserID,
		PracticumID: progress.PracticumID,
		Progress:    progress.Progress,
		CompletedAt: completedAt,
		LastUpdated: progress.LastUpdated.Format(time.RFC3339),
	}

	response.NewSuccessResponse(w, res, "User practicum progress retrieved successfully")
}

func (h *UserPracticumProgressHandler) UpdateProgress(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	var req dto.UpdateUserPracticumProgressRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	progress := model.UserPracticumProgress{
		ID:       id,
		Progress: req.Progress,
	}

	err = h.service.UpdateProgress(&progress)
	if err != nil {
		appErr := pkg.NewAppError("Failed to update progress", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "User practicum progress updated successfully")
}

func (h *UserPracticumProgressHandler) MarkAsCompleted(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid user ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	practicumID, err := strconv.Atoi(r.PathValue("practicum_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid practicum ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	err = h.service.MarkAsCompleted(userID, practicumID)
	if err != nil {
		appErr := pkg.NewAppError("Failed to mark progress as completed", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "User practicum progress marked as completed")
}

func (h *UserPracticumProgressHandler) DeleteProgress(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	err = h.service.DeleteProgress(id)
	if err != nil {
		appErr := pkg.NewAppError("Failed to delete progress", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "User practicum progress deleted successfully")
}
