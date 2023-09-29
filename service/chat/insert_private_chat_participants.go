package chatservice

import (
	"context"
	"github.com/google/uuid"
	"hangout/entity"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) InsertPrivateChatParticipants(ctx context.Context, req param.InsertPrivateChatParticipantsRequest) (*param.InsertPrivateChatParticipantsResponse, error) {
	const op = "ChatService.InsertPrivateChatParticipants"
	p1 := entity.PrivateChatParticipant{
		ID:     uuid.NewString(),
		ChatID: req.ChatID,
		UserID: req.UserID1,
	}
	p2 := entity.PrivateChatParticipant{
		ID:     uuid.NewString(),
		ChatID: req.ChatID,
		UserID: req.UserID2,
	}

	if err := s.chatRepo.AddToPrivateChatParticipants(ctx, p1, p2); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.InsertPrivateChatParticipantsResponse{}, nil
}
