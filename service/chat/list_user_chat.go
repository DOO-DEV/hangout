package chatservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) ListUserChats(ctx context.Context, _ param.GetUserChatsRequest, userID string) (*param.GetUserChatResponse, error) {
	const op = "ChatService.LIstUserChats"

	list, err := s.chatRepo.GetUserChatList(ctx, userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	chatIDs := make([]string, 0)
	for _, v := range list {
		chatIDs = append(chatIDs, v.ID)
	}
	return &param.GetUserChatResponse{Data: chatIDs}, nil
}
