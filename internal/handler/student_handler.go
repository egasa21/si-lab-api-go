package handler

import (
	"encoding/json"
	"fmt"

	"net/http"
	"strconv"

	"github.com/egasa21/si-lab-api-go/internal/dto"
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/pkg"
	"github.com/rs/zerolog/log"

	"github.com/egasa21/si-lab-api-go/internal/pkg/response"
	"github.com/egasa21/si-lab-api-go/internal/service"
)

type StudentHandler struct {
	studentService     service.StudentService
	studentDataService service.StudentDataService
}

func NewStudentHandler(studentService service.StudentService, studentDataService service.StudentDataService) *StudentHandler {
	return &StudentHandler{
		studentService:     studentService,
		studentDataService: studentDataService,
	}
}

func (h *StudentHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	// Get page and limit query parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1 // Default page number
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10 // Default limit
	}

	students, total, err := h.studentService.GetAllStudents(page, limit)
	if err != nil {
		// Return error response with status 500
		appErr := pkg.NewAppError("Unable to fetch students", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	// Calculate total pages for pagination
	totalPages := (total + limit - 1) / limit // Ceiling division

	pagination := response.Pagination{
		Page:       page,
		PerPage:    limit,
		TotalPages: totalPages,
		TotalItems: total,
	}

	response.NewPaginatedSuccessResponse(w, students, pagination, "Students retrieved successfully")
}

func (h *StudentHandler) GetStudentById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	student, err := h.studentService.GetStudentByID(id)
	if err != nil {

		appErr := pkg.NewAppError("Student not found", http.StatusNotFound)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, student, "Student retrieved successfully")
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {

	var student model.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {

		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	// check existing student
	existingStudent, err := h.studentService.GetStudentByStudentID(student.StudentIDNumber)
	if err == nil && existingStudent != nil {

		appErr := pkg.NewAppError("Student with this ID number is already registered", http.StatusConflict)
		response.NewErrorResponse(w, appErr)
		return
	}

	studentID, err := h.studentService.CreateStudent(&student)
	if err != nil {

		fmt.Print(err)
		appErr := pkg.NewAppError("Failed to create student", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	responseData := dto.StudentCreateResponse{
		StudentID: studentID,
	}

	response.NewSuccessResponse(w, responseData, "Student created successfully")
}

func (h *StudentHandler) GetStudentPracticumActivities(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	studentActivities, err := h.studentDataService.GetStudentPracticumActivity(id)
	if err != nil {
		log.Error().Err(err).Msg("Failed ngab")
		appErr := pkg.NewAppError("Student Activities not found", http.StatusNotFound)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, studentActivities, "student activities retrieved successfully")
}

func (h *StudentHandler) GetStudentSchedules(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("student_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	studentSchedules, err := h.studentDataService.GetStudentSchedules(id)
	if err != nil {
		log.Error().Err(err).Msg("Failed ngab")
		appErr := pkg.NewAppError("Student Activities not found", http.StatusNotFound)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, studentSchedules, "student schedules retrieved successfully")
}
