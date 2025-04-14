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

type PracticumModuleContentHandler struct {
	service service.PracticumModuleContentService
}

func NewPracticumModuleContentHandler(service service.PracticumModuleContentService) *PracticumModuleContentHandler {
	return &PracticumModuleContentHandler{service: service}
}

func (h *PracticumModuleContentHandler) CreateContent(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePracticumModuleContentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	content := model.PracticumModuleContent{
		IDModule: req.IDModule,
		Title:    req.Title,
		Content:  req.Content,
		Sequence: req.Sequence,
	}

	newModuleContent, err := h.service.CreateContent(&content)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create practicum module content")
		appErr := pkg.NewAppError("Failed to create content", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	moduleContentResponse := &dto.PracticumModuleContentResponse{
		ID:    uint(newModuleContent.IDContent),
		Title: newModuleContent.Title,
	}

	response.NewSuccessResponse(w, moduleContentResponse, "Content created successfully")
}

func (h *PracticumModuleContentHandler) GetContentByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	content, err := h.service.GetContentByID(id)
	if err != nil {
		appErr := pkg.NewAppError("Content not found", http.StatusNotFound)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, content, "Content retrieved successfully")
}

func (h *PracticumModuleContentHandler) GetContentsByModuleID(w http.ResponseWriter, r *http.Request) {
	moduleID, err := strconv.Atoi(r.PathValue("module_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid module ID", http.StatusBadRequest)
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

	contents, total, err := h.service.GetContentsByModuleID(moduleID, page, limit)
	if err != nil {
		appErr := pkg.NewAppError("Unable to fetch contents", http.StatusInternalServerError)
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

	response.NewPaginatedSuccessResponse(w, contents, pagination, "Contents retrieved successfully")
}

func (h *PracticumModuleContentHandler) UpdateContentByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	var req dto.UpdatePracticumModuleContentRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	updatedContent := model.PracticumModuleContent{
		IDModule:   req.IDModule,
		Title:      req.Title,
		Content:    req.Content,
		Sequence:   req.Sequence,
		MaterialID: req.MaterialID,
	}

	err = h.service.UpdateContentByID(id, &updatedContent)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update practicum module content")
		appErr := pkg.NewAppError("Failed to update content", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "Content updated successfully")
}
