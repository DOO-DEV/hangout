package password

import (
	"golang.org/x/crypto/bcrypt"
	"hangout/pkg/apperror"
)

func EncodePassword(passwordStr string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordStr), bcrypt.DefaultCost)
	if err != nil {
		return "", apperror.EncodePasswordErr
	}

	return string(hash), nil
}

func DecodePassword(rowPassword, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rowPassword)); err != nil {
		return apperror.DecodePasswordErr
	}

	return nil
}
