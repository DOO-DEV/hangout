package chatservice

import (
	"context"
	"hangout/entity"
)

type userRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
}

type chatRepository interface {
	GetUserChatList(ctx context.Context, userID string) ([]entity.Chat, error)
	GetChatMessages(ctx context.Context, chatID string) ([]entity.Message, error)
	SaveMessage(ctx context.Context, m entity.Message) error
	CreateChat(ctx context.Context, c *entity.Chat) (*entity.Chat, error)
	GetChatByUsersIds(ctx context.Context, u1, u2 string) (*entity.Chat, error)
}

type groupRepository interface {
	CheckUserGroupConnection(ctx context.Context, u1, u2 string) (bool, error)
}

type Service struct {
	chatRepo  chatRepository
	groupRepo groupRepository
	userRepo  userRepository
}

func New(chRepo chatRepository, gRepo groupRepository, uRepo userRepository) Service {
	return Service{
		chatRepo:  chRepo,
		groupRepo: gRepo,
		userRepo:  uRepo,
	}
}
