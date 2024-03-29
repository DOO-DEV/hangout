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
	AddToPrivateChatParticipants(ctx context.Context, p1, p2 entity.PrivateChatParticipant) error
	GetPrivateChatMessages(ctx context.Context, chatID string) ([]entity.Message, error)
}

type groupRepository interface {
	CheckUserGroupConnection(ctx context.Context, u1, u2 string) (bool, error)
}

type DeliveryStorage interface {
	PublishToPrivateMessage(message entity.Message, receiverID string) error
	SubscribeToPrivateMessage(userID string, msgChan chan<- entity.Message)
}

type Service struct {
	chatRepo        chatRepository
	groupRepo       groupRepository
	deliveryStorage DeliveryStorage
}

func New(chRepo chatRepository, gRepo groupRepository, deliveryStorage DeliveryStorage) Service {
	return Service{
		chatRepo:        chRepo,
		groupRepo:       gRepo,
		deliveryStorage: deliveryStorage,
	}
}
