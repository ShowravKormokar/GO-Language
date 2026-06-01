package middleware

import (
	"context"
	"encoding/json"
	"httpServer_JWT_MongoDB/dto"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserContextKey contextKey = "user"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		authHeader := rq.Header.Get("Authorization")

		if authHeader == "" {
			rw.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(rw).Encode(dto.BasicResponse{
				Success: false,
				Message: "Authorization header missing",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("my_super_secret_256key"), nil
		})
		if err != nil || !token.Valid {

			rw.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(rw).Encode(dto.BasicResponse{
				Success: false,
				Message: "Invalid token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			rw.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(rw).Encode(
				dto.BasicResponse{
					Success: false,
					Message: "Invalid token claims",
				},
			)
			return
		}

		ctx := context.WithValue(rq.Context(), UserContextKey, claims)
		next.ServeHTTP(rw, rq.WithContext(ctx))
	})
}
