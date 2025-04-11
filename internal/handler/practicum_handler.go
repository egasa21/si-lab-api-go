package handler

import (
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/pkg"
	"github.com/egasa21/si-lab-api-go/internal/pkg/response"
	"github.com/egasa21/si-lab-api-go/internal/service"
	"github.com/rs/zerolog/log"
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

func (h *PracticumHandler) GetPracticumByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	practicum, err := h.service.GetPracticumByID(id)
	if err != nil {
		appErr := pkg.NewAppError("Practicum not found", http.StatusNotFound)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, practicum, "practicum retrieved successfully")
}

func (h *PracticumHandler) GetAllPracticums(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	practicums, total, err := h.service.GetAllPracticums(page, limit)
	if err != nil {

		appErr := pkg.NewAppError("Unable to fetch practicums", http.StatusInternalServerError)
		log.Error().Err(err)
		response.NewErrorResponse(w, appErr)
		return
	}

	totalPages := (total + limit - 1) / limit

	pagination := response.Pagination{
		Page:       page,
		PerPage:    limit,
		TotalPages: totalPages,
		TotalItems: total,
	}

	response.NewPaginatedSuccessResponse(w, practicums, pagination, "Students retrieved successfully")
}

func (h *PracticumHandler) GetPracticumWithMaterialContents(w http.ResponseWriter, r *http.Request) {
	practicumID, err := strconv.Atoi(r.PathValue("practicum_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid practicum ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	practicumWithMaterials, err := h.service.GetPracticumWithMaterialContents(practicumID)
	if err != nil {
		appErr := pkg.NewAppError("practicum not found", http.StatusNotFound)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, practicumWithMaterials, "Practicum with materials retrieved successfully")
}
