package chatservice

import (
	"context"
	"github.com/google/uuid"
	"hangout/entity"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) CreateGroupChat(ctx context.Context, req param.CreateGroupChatRequest) (*param.CreateGroupChatResponse, error) {
	const op = "ChatService.CreateGroupChat"

	c := entity.GroupChat{
		ID:   uuid.NewString(),
		Name: req.Name,
	}
	groupChat, err := s.chatRepo.CreateGroupChat(ctx, c)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.CreateGroupChatResponse{
		Name: groupChat.Name,
		ID:   groupChat.ID,
	}, nil
}
