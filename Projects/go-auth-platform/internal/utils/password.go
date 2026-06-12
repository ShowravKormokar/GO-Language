package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	BcryptCost = bcrypt.DefaultCost
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		BcryptCost,
	)

	if err != nil {
		return "", err
	}

	return string(bytes), err
}

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
