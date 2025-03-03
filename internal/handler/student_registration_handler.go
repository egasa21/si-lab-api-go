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

type StudentRegistrationHandler struct {
	service service.StudentRegistrationService
}

func NewStudentRegistrationHandler(service service.StudentRegistrationService) *StudentRegistrationHandler {
	return &StudentRegistrationHandler{service: service}
}

// RegisterStudent handles the creation of student registration
func (h *StudentRegistrationHandler) RegisterStudent(w http.ResponseWriter, r *http.Request) {
	var req dto.StudentRegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	// Loop through the PracticumIDs to register each practicum
	for _, practicumID := range req.PracticumIDs {
		registration := model.StudentRegistration{
			StudentID:   req.StudentID,
			PracticumID: practicumID,
		}
		err := h.service.RegisterStudent(&registration)
		if err != nil {
			appErr := pkg.NewAppError("Failed to register student for practicum", http.StatusInternalServerError)
			response.NewErrorResponse(w, appErr)
			return
		}
	}

	response.NewSuccessResponse(w, nil, "Student registration successful")
}

// GetRegistrationsByStudentID retrieves all registrations for a student
func (h *StudentRegistrationHandler) GetRegistrationsByStudentID(w http.ResponseWriter, r *http.Request) {
	studentID, err := strconv.Atoi(r.PathValue("student_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid student ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	registrations, err := h.service.GetRegistrationsByStudentID(studentID)
	if err != nil {
		appErr := pkg.NewAppError("Unable to fetch registrations", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, registrations, "Registrations retrieved successfully")
}

// GetRegistrationsByPracticumID retrieves all registrations for a practicum
func (h *StudentRegistrationHandler) GetRegistrationsByPracticumID(w http.ResponseWriter, r *http.Request) {
	practicumID, err := strconv.Atoi(r.URL.Query().Get("practicum_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid practicum ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	registrations, err := h.service.GetRegistrationsByPracticumID(practicumID)
	if err != nil {
		appErr := pkg.NewAppError("Unable to fetch registrations", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, registrations, "Registrations retrieved successfully")
}

// DeleteRegistration handles the deletion of a student registration by its ID
func (h *StudentRegistrationHandler) DeleteRegistration(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid registration ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	err = h.service.DeleteRegistration(id)
	if err != nil {
		appErr := pkg.NewAppError("Failed to delete registration", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "Registration deleted successfully")
}
