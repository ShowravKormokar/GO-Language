package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateResetToken() string {
	bytes := make([]byte, 32)

	// Read random byte
	rand.Read(bytes)

	return hex.EncodeToString(bytes)
}

// Hash the reset token to store on db
func HashResetToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
