package chatservice

import (
	"context"
	"github.com/google/uuid"
	"hangout/entity"
	param "hangout/param/http"
	"hangout/pkg/richerror"
	"slices"
	"strings"
)

func (s Service) Create(ctx context.Context, req param.CreatePrivateChatRequest) (*param.CreatePrivateChatResponse, error) {
	const op = "ChatService.Create"

	// the name of private chats are "userid-userid" and sort from higher-lower
	users := []string{req.Receiver, req.Sender}
	slices.Sort(users)

	chatName := strings.Join(users, "-")

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
