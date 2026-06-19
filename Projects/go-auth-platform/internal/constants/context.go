package constants

type ContextKey string

const (
	ContextUserID ContextKey = "user_id"
	ContextRole   ContextKey = "role"
	ContextClaims ContextKey = "claims"
)
