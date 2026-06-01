package middleware

import (
	"encoding/json"
	"httpServer_JWT_MongoDB/dto"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {

		claims, ok := rq.Context().Value(ClaimsKey).(jwt.MapClaims)
		if !ok || claims["role"].(string) != "admin" {
			rw.WriteHeader(http.StatusForbidden)
			json.NewEncoder(rw).Encode(dto.BasicResponse{
				Success: false,
				Message: "Admin access required",
			})

			return
		}

		next.ServeHTTP(rw, rq)
	})
}
