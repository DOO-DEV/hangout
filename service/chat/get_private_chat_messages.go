package chatservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) GetPrivateChatMessages(ctx context.Context, req param.GetPrivateChatMessages) (*param.GetPrivateChatMessagesRepose, error) {
	const op = "ChatService.GetPrivateChatMessages"

	msgs, err := s.chatRepo.GetPrivateChatMessages(ctx, req.ChatID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	me := make([]param.PrivateChatMessages, 0)
	for _, m := range msgs {
		r := param.PrivateChatMessages{
			ID:        m.ID,
			Timestamp: m.Timestamp,
			Content:   m.Content,
			Status:    int(m.Status),
		}

		me = append(me, r)
	}

	return &param.GetPrivateChatMessagesRepose{Data: me}, nil
}
