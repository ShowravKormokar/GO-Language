package utils

import (
	"go-auth-platform/internal/config"
	dto "go-auth-platform/internal/dto/claims"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenPair struct {
	AccessToken, RefreshToken string
}

// Access token
func GenerateAccessToken(userID string, email string, role string) (string, string, error) {
	jti := uuid.NewString()

	claims := dto.JWTClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		JTI:    jti,

		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ID:        jti,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.AppConfig.JWTAccessTTL)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		return "", "", err
	}

	return signed, jti, nil
}

// Refresh token
func GenerateRefreshToken(userID string) (string, error) {
	jti := uuid.NewString()

	claims := jwt.RegisteredClaims{
		Subject: userID,
		ID:      jti,

		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.AppConfig.JWTRefreshTTL)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}
