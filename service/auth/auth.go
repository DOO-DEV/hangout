package authservice

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"hangout/entity"
	"time"
)

type Config struct {
	SignKey                  string `koanf:"sign_key"`
	Prefix                   string `koanf:"prefix"`
	AuthMiddlewareContextKey string `koanf:"auth_middleware_context_key"`
}

type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{config: cfg}
}

func (s Service) CreateToken(u *entity.User) (string, error) {
	claims := Claims{
		Username: u.Username,
		ID:       u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    u.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.SignKey))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}

func (s Service) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
