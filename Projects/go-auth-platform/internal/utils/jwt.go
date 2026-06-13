package utils

import (
	"errors"
	"go-auth-platform/internal/config"
	dto "go-auth-platform/internal/dto/claims"
	"net/http"
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

// Token pair
func GenerateTokenPair(userID, email, role string) (*TokenPair, string, error) {
	accessToken, jti, err := GenerateAccessToken(userID, email, role)
	if err != nil {
		return nil, "", err
	}

	refreshToken, err := GenerateRefreshToken(userID)
	if err != nil {
		return nil, "", err
	}

	return &TokenPair{AccessToken: accessToken, RefreshToken: refreshToken}, jti, nil
}

// Parse access token
func ParseAccessToken(tokenString string) (*dto.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&dto.JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("invalid signing method")
			}

			return []byte(config.AppConfig.JWTSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*dto.JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// Parse refresh token
func ParseRefreshToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JWTSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	return claims, nil
}

// Cookie helper to set access token on cookie
func SetAccessCookie(rw http.ResponseWriter, token string) {
	http.SetCookie(
		rw,
		&http.Cookie{
			Name:     "access_token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   config.AppConfig.CookieSecure,
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now().Add(config.AppConfig.JWTAccessTTL),
		},
	)
}

// Set refresh token on cookie
func SetRefreshCookie(rw http.ResponseWriter, token string) {
	http.SetCookie(
		rw,
		&http.Cookie{
			Name:     "refresh_token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   config.AppConfig.CookieSecure,
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now().Add(config.AppConfig.JWTRefreshTTL),
		},
	)
}

func ClearAuthCookies(w http.ResponseWriter) {

	http.SetCookie(
		w,
		&http.Cookie{
			Name:     "access_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			MaxAge:   -1,
		},
	)

	http.SetCookie(
		w,
		&http.Cookie{
			Name:     "refresh_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			MaxAge:   -1,
		},
	)
}
