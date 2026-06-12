package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	BcryptCost = bcrypt.DefaultCost
)

// Convert password into hash from
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		BcryptCost,
	)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Check password match or not
func CheckPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)

	if err != nil {
		return errors.New("invalid credentials")
	}

	return nil
}
