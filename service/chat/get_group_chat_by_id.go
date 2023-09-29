package chatservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) GetGroupChatByID(ctx context.Context, req param.GetGroupChatByIDRequest) (*param.GetGroupChatByIDResponse, error) {
	const op = "ChatService.GetGroupChatByID"

	chat, err := s.chatRepo.GetGroupChatByID(ctx, req.ChatID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}
	if chat == nil {
		return nil, nil
	}

	return &param.GetGroupChatByIDResponse{
		ChatID:   chat.ID,
		ChatName: chat.Name,
	}, nil
}
