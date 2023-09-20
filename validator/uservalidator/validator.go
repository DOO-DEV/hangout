package uservalidator

import (
	"context"
)

type repository interface {
	IsUserExists(ctx context.Context, username string) (bool, error)
}

type Validator struct {
	repo repository
}

func New(repo repository) Validator {
	return Validator{
		repo: repo,
	}
}
