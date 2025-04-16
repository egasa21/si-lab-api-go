package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/egasa21/si-lab-api-go/internal/dto"
	"github.com/egasa21/si-lab-api-go/internal/middlewares"
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/pkg"
	"github.com/egasa21/si-lab-api-go/internal/pkg/response"
	"github.com/egasa21/si-lab-api-go/internal/service"
	"github.com/rs/zerolog/log"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest

	// Decode the JSON payload into the request struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	// Log or check if the password exists
	fmt.Printf("Password received in request: %s\n", req.Password)

	// Populate the User model using the request data
	user := model.User{
		Email:     req.Email,
		Password:  req.Password, // Pass the raw password to be hashed later
		IDStudent: req.IDStudent,
	}

	// Call the service with the User model and default roles
	if err := h.authService.Register(&user, nil); err != nil {
		appErr := pkg.NewAppError(err.Error(), http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	// Respond with success
	response.NewSuccessResponse(w, nil, "User registered successfully")
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

	data := map[string]interface{}{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"expiresIn":     token.ExpiresIn,
	}

	response.NewSuccessResponse(w, data, "login successfully ")

}

func (h *AuthHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	userIDValue := r.Context().Value(middlewares.UserIDKey)
	if userIDValue == nil {
		log.Error().Msg("User ID not found in context")
		response.NewErrorResponse(w, pkg.NewAppError("Unauthorized", http.StatusUnauthorized))
		return
	}

	userIDFloat, ok := userIDValue.(float64)
	if !ok {
		log.Error().Msgf("Invalid user ID type in context: %T", userIDValue)
		response.NewErrorResponse(w, pkg.NewAppError("Unauthorized", http.StatusUnauthorized))
		return
	}
	userID := int(userIDFloat)

	userData, err := h.authService.GetUserByID(userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed ngab")
		appErr := pkg.NewAppError("user not found", http.StatusNotFound)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, userData, "user data retrieved successfully")
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.RefreshToken == "" {
		appErr := pkg.NewAppError("Invalid refresh token payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	newTokens, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		appErr := pkg.NewAppError("Failed to refresh token", http.StatusUnauthorized)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, map[string]interface{}{
		"access_token":  newTokens.AccessToken,
		"refresh_token": newTokens.RefreshToken,
		"expiresIn":     newTokens.ExpiresIn,
	}, "Token refreshed successfully")
}
