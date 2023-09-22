package chatservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) ListUserChats(ctx context.Context, req param.GetUserChatsRequest, userID string) (*param.GetUserChatResponse, error) {
	const op = "ChatService.LIstUserChats"

	list, err := s.chatRepo.GetUserChatList(ctx, userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.GetUserChatResponse{Data: list}, nil
}
