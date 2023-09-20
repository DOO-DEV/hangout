package groupservice

import (
	"context"
	"hangout/entity"
)

type Repository interface {
	CreateGroup(ctx context.Context, group entity.Group) error
	GetAllGroups(ctx context.Context) ([]*entity.Group, error)
	GetMyGroup(ctx context.Context, userID string) ([]entity.Member, error)
	CheckUserGroup(ctx context.Context, username string) (bool, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
