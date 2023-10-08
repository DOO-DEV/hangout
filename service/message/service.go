package messageservice

import (
	"context"
	"hangout/entity"
)

type MessageRepository interface {
	SavePrivateMessage(ctx context.Context, msg entity.Message) (*entity.Message, error)
	SaveGroupMessage(ctx context.Context, msg entity.Message) (*entity.Message, error)
	SetPrivateMessageAsRead(ctx context.Context, messageID string) error
}
type Service struct {
	messageRepo MessageRepository
}

func New(msgRepo MessageRepository) Service {
	return Service{
		messageRepo: msgRepo,
	}
}
