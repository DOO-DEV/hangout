package userservice

import (
	"context"
	"hangout/entity"
)

type Repository interface {
	Register(ctx context.Context, user *entity.User) error
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type AuthGenerator interface {
	CreateToken(u *entity.User) (string, error)
}

type Service struct {
	repo          Repository
	authGenerator AuthGenerator
}

func New(repo Repository, authGen AuthGenerator) Service {
	return Service{
		repo:          repo,
		authGenerator: authGen,
	}
}
