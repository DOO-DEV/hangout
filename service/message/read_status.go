package messageservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) SetStatusRead(ctx context.Context, req param.SetStatusReadRequest) (*param.SetStatusReadResponse, error) {
	const op = "MessageService.SetStatusRead"

	if err := s.messageRepo.SetPrivateMessageAsRead(ctx, req.MessageID); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.SetStatusReadResponse{}, nil
}
