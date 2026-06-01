package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID, userEmail, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"u_email": userEmail,
		"role":    role,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte("my_super_secret_256key"))
}
