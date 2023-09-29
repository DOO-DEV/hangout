package chatservice

import (
	"context"
	"github.com/google/uuid"
	"hangout/entity"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) Create(ctx context.Context, req param.CreateChatRequest) (*param.CreateChatResponse, error) {
	const op = "ChatService.Create"

	c := entity.Chat{
		ID:   uuid.NewString(),
		Type: entity.ChatType(req.Type),
		Name: req.Name,
	}
	newChat, err := s.chatRepo.CreateChat(ctx, &c)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.CreateChatResponse{
		ChatID:   newChat.ID,
		ChatName: newChat.Name,
	}, nil
}
