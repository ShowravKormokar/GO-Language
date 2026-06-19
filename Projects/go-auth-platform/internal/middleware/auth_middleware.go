package middleware

import (
	"context"
	"go-auth-platform/internal/constants"
	"go-auth-platform/internal/repository"
	"go-auth-platform/internal/utils"
	"net/http"
)

func AuthRequired(blacklistRepo repository.BlacklistRepository) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(

			func(rw http.ResponseWriter, req *http.Request) {

				cookie, err := req.Cookie("access_token")
				if err != nil {
					http.Error(
						rw,
						"authentication required",
						http.StatusUnauthorized,
					)
					return
				}

				claims, err := utils.ParseAccessToken(cookie.Value)
				if err != nil{
					http.Error(
						rw,
						"invalid token",
						http.StatusUnauthorized,
					)
					return
				}

				exists, err := blacklistRepo.ExistsByJTI(req.Context(), claims.JTI)
				if err != nil {
					http.Error(
						rw,
						"server error",
						http.StatusInternalServerError,
					)
					return
				}

				if exists{
					http.Error(
						rw,
						"token revoked",
						http.StatusUnauthorized,
					)
					return
				}

				ctx := context.WithValue(req.Context(),constants.ContextClaims, claims)
				ctx = context.WithValue(ctx,constants.ContextUserID,claims.UserID)
				ctx = context.WithValue(ctx,constants.ContextRole,claims.Role)

				next.ServeHTTP(
					rw,
					req.WithContext(ctx),
				)
			},
		)
	}
}