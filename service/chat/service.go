package chatservice

import (
	"context"
	"hangout/entity"
	dbparam "hangout/param/pgdb"
)

type chatRepository interface {
	SaveMessage(ctx context.Context, m entity.Message) error
	GetChatMessages(ctx context.Context, sender, receiver string) ([]entity.Message, error)
	GetUserChatList(ctx context.Context, userID string) ([]dbparam.Chat, error)
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
