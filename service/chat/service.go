package chatservice

import (
	"context"
	"hangout/entity"
)

type chatRepository interface {
	SaveMessage(ctx context.Context, m entity.Message) error
}

type groupRepository interface {
	CheckUserGroupConnection(ctx context.Context, u1, u2 string) (bool, error)
}

type Service struct {
	chatRepo  chatRepository
	groupRepo groupRepository
}

func New(chRepo chatRepository, gRepo groupRepository) Service {
	return Service{
		chatRepo:  chRepo,
		groupRepo: gRepo,
	}
}
