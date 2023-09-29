package messageservice

import (
	"context"
	"github.com/google/uuid"
	"hangout/entity"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) SavePrivateMessage(ctx context.Context, req param.PrivateMessageRequest) (*param.PrivateMessageResponse, error) {
	const op = "MessageService.SavePrivateMessages"

	m := entity.Message{
		ID:       uuid.NewString(),
		ChatID:   req.ChatID,
		SenderID: req.SenderID,
		Content:  req.Content,
		Type:     entity.MsgType(req.Type),
		Status:   entity.MsgStatusDelivered,
	}
	msg, err := s.messageRepo.SavePrivateMessage(ctx, m)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.PrivateMessageResponse{
		Timestamp: msg.Timestamp,
		ID:        msg.ID,
	}, nil
}
