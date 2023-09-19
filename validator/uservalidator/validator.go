package uservalidator

import (
	"context"
	"hangout/entity"
)

type repository interface {
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type Validator struct {
	repo repository
}

func New(repo repository) Validator {
	return Validator{
		repo: repo,
	}
}
