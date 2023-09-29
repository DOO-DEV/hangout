package chatservice

import (
	"context"
	"hangout/entity"
)

type chatRepository interface {
	CreatePrivateChat(ctx context.Context, chat entity.PrivateChat) (*entity.PrivateChat, error)
	GetPrivateChatByName(ctx context.Context, name string) (*entity.PrivateChat, error)
	CreateGroupChat(ctx context.Context, group entity.GroupChat) (*entity.GroupChat, error)
	GetGroupChatByID(ctx context.Context, groupID string) (*entity.GroupChat, error)
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
