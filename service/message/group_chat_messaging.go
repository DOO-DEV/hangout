package messageservice

import (
	"context"
	"github.com/google/uuid"
	"hangout/entity"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) SaveGroupChatMessage(ctx context.Context, req param.GroupChatMessageRequest) (*param.GroupChatMessageResponse, error) {
	const op = "MessageService.SaveGroupChatMessage"

	m := entity.Message{
		ID:       uuid.NewString(),
		ChatID:   req.ChatID,
		SenderID: req.SenderID,
		Content:  req.Content,
		Type:     entity.MsgType(req.Type),
		Status:   entity.MsgStatusDelivered,
	}
	msg, err := s.messageRepo.SaveGroupMessage(ctx, m)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.GroupChatMessageResponse{
		Timestamp: msg.Timestamp,
		ID:        msg.ID,
	}, nil
}
