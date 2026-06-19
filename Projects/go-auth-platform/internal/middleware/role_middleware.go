package middleware

import (
	"go-auth-platform/internal/constants"
	"net/http"
)

func RequireMinRole(required string) func(http.Handler) http.Handler {

	// This middleware should be used after AuthRequired, as it relies on the role being set in the context
	return func(next http.Handler) http.Handler {

		// This middleware checks if the user's role meets the minimum required role for the endpoint
		return http.HandlerFunc(
			
			// It retrieves the user's role from the context and compares it against the required role using a predefined hierarchy
			func(rw http.ResponseWriter, req *http.Request) {
				// If the user's role does not meet the required level, it responds with a 403 Forbidden status
				role, ok := req.Context().Value(constants.ContextRole).(string)
				if !ok {
					http.Error(
						rw,
						"forbidden",
						http.StatusForbidden,
					)
					return
				}
				// The role hierarchy is defined in constants.RoleHierarchy, where higher values represent higher privileges
				current := constants.RoleHierarchy[role]

				// The required role's hierarchy level is retrieved, and if the user's current role level is less than the required level, access is denied
				requiredLevel := constants.RoleHierarchy[required]
				if current < requiredLevel {
					http.Error(
						rw,
						"forbidden",
						http.StatusForbidden,
					)
					return
				}

				// If the user's role meets the requirement, the request is passed to the next handler in the chain
				next.ServeHTTP(
					rw,
					req,
				)
			},
		)
	}
}
