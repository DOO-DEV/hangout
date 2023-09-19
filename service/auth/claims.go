package authservice

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string
	ID       string
	jwt.RegisteredClaims
}

func (c Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}
