package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/egasa21/si-lab-api-go/internal/dto"
	"github.com/egasa21/si-lab-api-go/internal/pkg"
	"github.com/egasa21/si-lab-api-go/internal/pkg/response"
	"github.com/egasa21/si-lab-api-go/internal/service"
)

type StudentClassEnrollmentHandler struct {
	service service.StudentClassEnrollmentService
}

// NewStudentClassEnrollmentHandler initializes a new handler
func NewStudentClassEnrollmentHandler(service service.StudentClassEnrollmentService) *StudentClassEnrollmentHandler {
	return &StudentClassEnrollmentHandler{service: service}
}

// EnrollStudent handles enrolling a student into a class
func (h *StudentClassEnrollmentHandler) EnrollStudent(w http.ResponseWriter, r *http.Request) {
	var req dto.StudentEnrollmentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	err = h.service.EnrollStudent(req.ClassID, req.StudentID)
	if err != nil {
		appErr := pkg.NewAppError("Failed to enroll student", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "Student enrolled successfully")
}

// GetEnrollmentsByStudentID retrieves all class enrollments for a specific student
func (h *StudentClassEnrollmentHandler) GetEnrollmentsByStudentID(w http.ResponseWriter, r *http.Request) {
	studentID, err := strconv.Atoi(r.PathValue("student_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid student ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	enrollments, err := h.service.GetEnrollmentsByStudentID(studentID)
	if err != nil {
		appErr := pkg.NewAppError("Unable to fetch enrollments", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, enrollments, "Enrollments retrieved successfully")
}

// GetEnrollmentsByClassID retrieves all students enrolled in a specific class
func (h *StudentClassEnrollmentHandler) GetEnrollmentsByClassID(w http.ResponseWriter, r *http.Request) {
	classID, err := strconv.Atoi(r.URL.Query().Get("class_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid class ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	enrollments, err := h.service.GetEnrollmentsByClassID(classID)
	if err != nil {
		appErr := pkg.NewAppError("Unable to fetch enrollments", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, enrollments, "Class enrollments retrieved successfully")
}

// UnenrollStudent handles removing a student from a class
func (h *StudentClassEnrollmentHandler) UnenrollStudent(w http.ResponseWriter, r *http.Request) {
	enrollmentID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid enrollment ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	err = h.service.UnenrollStudent(enrollmentID)
	if err != nil {
		appErr := pkg.NewAppError("Failed to unenroll student", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "Student unenrolled successfully")
}
