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

type PracticumClassHandler struct {
	service service.PracticumClassService
}

func NewPracticumClassHandler(service service.PracticumClassService) *PracticumClassHandler {
	return &PracticumClassHandler{service: service}
}

func (h *PracticumClassHandler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePracticumClassRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	class := model.PracticumClass{
		PracticumID: req.PracticumID,
		Name:        req.Name,
		Quota:       req.Quota,
		Day:         req.Day,
		Time:        req.Time,
	}

	err = h.service.CreateClass(&class)
	if err != nil {
		appErr := pkg.NewAppError("Failed to create practicum class", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "Practicum class created successfully")
}

func (h *PracticumClassHandler) GetClassByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	class, err := h.service.GetClassByID(id)
	if err != nil {
		appErr := pkg.NewAppError("Class not found", http.StatusNotFound)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, class, "Class retrieved successfully")
}

func (h *PracticumClassHandler) GetClassesByPracticumID(w http.ResponseWriter, r *http.Request) {
	practicumID, err := strconv.Atoi(r.PathValue("practicum_id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid practicum ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	classes, err := h.service.GetClassesByPracticumID(practicumID)
	if err != nil {
		appErr := pkg.NewAppError("Unable to fetch classes", http.StatusInternalServerError)
		log.Error().Err(err)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, classes, "Classes retrieved successfully")
}

func (h *PracticumClassHandler) UpdateClass(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	var req dto.UpdatePracticumClassRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	class := model.PracticumClass{
		IDPracticumClass: id,
		PracticumID:      req.PracticumID,
		Name:             req.Name,
		Quota:            req.Quota,
		Day:              req.Day,
		Time:             req.Time,
	}

	err = h.service.UpdateClass(&class)
	if err != nil {
		appErr := pkg.NewAppError("Failed to update practicum class", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "Practicum class updated successfully")
}

func (h *PracticumClassHandler) DeleteClass(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		appErr := pkg.NewAppError("Invalid ID", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	err = h.service.DeleteClass(id)
	if err != nil {
		appErr := pkg.NewAppError("Failed to delete practicum class", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "Practicum class deleted successfully")
}
