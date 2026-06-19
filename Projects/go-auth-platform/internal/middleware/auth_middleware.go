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

				// The middleware checks for the presence of an "access_token" cookie in the incoming HTTP request. If the cookie is not found, it responds with a 401 Unauthorized status, indicating that authentication is required.
				cookie, err := req.Cookie("access_token")
				if err != nil {
					http.Error(
						rw,
						"authentication required",
						http.StatusUnauthorized,
					)
					return
				}

				// If the cookie is present, the middleware attempts to parse and validate the access token using a utility function. If the token is invalid or cannot be parsed, it responds with a 401 Unauthorized status.
				claims, err := utils.ParseAccessToken(cookie.Value)
				if err != nil {
					http.Error(
						rw,
						"invalid token",
						http.StatusUnauthorized,
					)
					return
				}

				// The middleware then checks if the token's JTI (JWT ID) is present in a blacklist repository, which indicates that the token has been revoked. If the token is found in the blacklist, it responds with a 401 Unauthorized status, indicating that the token is no longer valid.
				exists, err := blacklistRepo.ExistsByJTI(req.Context(), claims.JTI)
				if err != nil {
					http.Error(
						rw,
						"server error",
						http.StatusInternalServerError,
					)
					return
				}

				if exists {
					http.Error(
						rw,
						"token revoked",
						http.StatusUnauthorized,
					)
					return
				}

				// short comment: if token is valid and not blacklisted, the middleware adds the token's claims, user ID, and role to the request context, allowing subsequent handlers to access this information for authorization purposes. Finally, it calls the next handler in the chain with the modified request context.
				ctx := context.WithValue(req.Context(), constants.ContextClaims, claims)
				ctx = context.WithValue(ctx, constants.ContextUserID, claims.UserID)
				ctx = context.WithValue(ctx, constants.ContextRole, claims.Role)

				next.ServeHTTP(
					rw,
					req.WithContext(ctx),
				)
			},
		)
	}
}
