package auth

import (
	"time"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super_secfet_ket")

type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// GenerateJWT generates an access token and a refresh token for the user
func GenerateJWT(userID int, roles []model.RoleModel) (*TokenDetails, error) {
	// Convert roles to string slice
	roleNames := make([]string, len(roles))
	for i, role := range roles {
		roleNames[i] = string(role.Name)
	}

	// Access token claims
	accessTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"roles":   roleNames,
		"exp":     time.Now().Add(10 * time.Minute).Unix(),
	}

	// Refresh token claims
	refreshTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// Return tokens
	return &TokenDetails{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
