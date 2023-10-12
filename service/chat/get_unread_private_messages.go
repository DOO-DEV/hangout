package chatservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) GetUnreadPrivateChatMessages(ctx context.Context, req param.GetUnreadPrivateChatMessagesRequest) (*param.GetUnreadPrivateChatMessagesResponse, error) {
	const op = "ChatService.GetUnreadPrivateChatMessages"

	messages, err := s.chatRepo.GetPrivateChatMessages(ctx, req.ChatID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	me := make([]param.PrivateChatMessages, 0)
	for _, m := range messages {
		r := param.PrivateChatMessages{
			ID:        m.ID,
			Timestamp: m.Timestamp,
			Content:   m.Content,
			Status:    int(m.Status),
		}

		me = append(me, r)
	}

	return &param.GetUnreadPrivateChatMessagesResponse{Data: me}, nil
}
