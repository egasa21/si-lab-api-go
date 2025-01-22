package middlewares

import (
	"context"
	"net/http"

	"github.com/egasa21/si-lab-api-go/internal/pkg"
	"github.com/egasa21/si-lab-api-go/internal/pkg/response"
	"github.com/egasa21/si-lab-api-go/internal/service"
	"github.com/egasa21/si-lab-api-go/internal/utils"
	"github.com/egasa21/si-lab-api-go/pkg/auth"
)

const UserIDKey utils.ContextKey = "user_id"

func AuthMiddleware(authService service.AuthService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
				response.NewErrorResponse(w, &pkg.AppError{
					Message:    "Missing Authorization header",
					StatusCode: http.StatusUnauthorized,
				})
				return
			}

			tokenString := authHeader[7:]
			claims, err := auth.VerifyToken(tokenString)
			if err != nil {
				response.NewErrorResponse(w, &pkg.AppError{
					Message:    "Invalid or expired token",
					StatusCode: http.StatusUnauthorized,
				})
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims["user_id"])
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
