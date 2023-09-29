package chatservice

import (
	"context"
	"github.com/google/uuid"
	"hangout/entity"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) Create(ctx context.Context, req param.CreatePrivateChatRequest) (*param.CreatePrivateChatResponse, error) {
	const op = "ChatService.Create"

	chatName := s.createPrivateChatName(req.Sender, req.Receiver)

	c := entity.PrivateChat{
		ID:   uuid.NewString(),
		Name: chatName,
	}
	newChat, err := s.chatRepo.CreatePrivateChat(ctx, c)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.CreatePrivateChatResponse{
		ChatID:   newChat.ID,
		ChatName: newChat.Name,
	}, nil
}
