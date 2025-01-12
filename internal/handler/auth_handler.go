package handler

import (
	"encoding/json"
	"net/http"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/pkg"
	"github.com/egasa21/si-lab-api-go/internal/pkg/response"
	"github.com/egasa21/si-lab-api-go/internal/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {

		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	roleNames := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roleNames[i] = string(role.Name)
	}

	if err := h.authService.Register(&user, nil); err != nil {
		appErr := pkg.NewAppError(err.Error(), http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)

		return
	}

	response.NewSuccessResponse(w, nil, "User registered successfully ")
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	token, err := h.authService.Login(request.Email, request.Password)
	if err != nil {
		appErr := pkg.NewAppError(err.Error(), http.StatusUnauthorized)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, token, "User registered successfully ")

}
