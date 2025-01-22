package auth

import (
	"errors"
	"os"
	"time"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var ErrInvalidToken = errors.New("invalid token")
var ErrTokenExpired = errors.New("token has expired")

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

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return jwtSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
