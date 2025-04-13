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
	"github.com/rs/zerolog/log"
)

type PracticumModuleHandler struct {
	service service.PracticumModuleService
}

func NewPracticumModuleHandler(service service.PracticumModuleService) *PracticumModuleHandler {
	return &PracticumModuleHandler{service: service}
}

func (h *PracticumModuleHandler) CreateModule(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePracticumModuleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	module := model.PracticumModule{
		Title:       req.Title,
		PracticumID: req.PracticumID,
	}

	newModule, err := h.service.CreateModule(&module)
	if err != nil {
		appErr := pkg.NewAppError("Failed to create module", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	moduleResponse := &dto.PracticumModuleResponse{
		ID:    uint(newModule.ID),
		Title: newModule.Title,
	}

	response.NewSuccessResponse(w, moduleResponse, "Module created successfully")
}

func (h *PracticumModuleHandler) GetModuleByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	module, err := h.service.GetModuleByID(id)
	if err != nil {
		appErr := pkg.NewAppError("Module not found", http.StatusNotFound)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, module, "Module retrieved successfully")
}

func (h *PracticumModuleHandler) GetModulesByPracticumID(w http.ResponseWriter, r *http.Request) {
	practicumID, err := strconv.Atoi(r.PathValue("practicum_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid practicum ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

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

	modules, total, err := h.service.GetModulesByPracticumID(practicumID, page, limit)
	if err != nil {
		appErr := pkg.NewAppError("Unable to fetch modules", http.StatusInternalServerError)
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

	response.NewPaginatedSuccessResponse(w, modules, pagination, "Modules retrieved successfully")
}
