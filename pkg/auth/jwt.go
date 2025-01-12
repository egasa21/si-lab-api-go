package auth

import (
	"time"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super_secfet_ket")

func GenerateJWT(userID int, roles []model.RoleModel) (string, error) {
	roleNames := make([]string, len(roles))
	for i, role := range roles {
		roleNames[i] = string(role.Name)
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"roles":   roleNames,
		"exp":     time.Now().Add(10 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
