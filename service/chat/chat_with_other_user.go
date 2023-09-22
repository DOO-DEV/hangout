package chatservice

import (
	"context"
	"errors"
	"hangout/entity"
	param "hangout/param/http"
	"hangout/pkg/errmsg"
	"hangout/pkg/richerror"
)

func (s Service) ChatWithOtherUser(ctx context.Context, req param.ChatMessageRequest, sender, receiver string) (*param.ChatMessageResponse, error) {
	const op = "ChatService.ChatWithOtherUser"

	// user can chat with user in two case:
	// 1- they are in same group (groupIDUser1 == groupIDUser2)
	// 2- they group are connected before (accept is true)
	// then add message to storage

	isConnect, err := s.groupRepo.CheckUserGroupConnection(ctx, sender, receiver)
	if !isConnect {
		if err != nil {
			return nil, richerror.New(op).WithError(err)
		}
		wErr := richerror.New(op).WithError(errors.New("")).WithMessage(errmsg.ErrorMsgUsersAreNotConnected).WithKind(richerror.KindForbidden)
		return nil, richerror.New(op).WithError(wErr)
	}
	m := entity.Message{
		Sender:   sender,
		Receiver: receiver,
		Content:  req.Content,
		Type:     entity.MsgTypeText,
	}
	if err := s.chatRepo.SaveMessage(ctx, m); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.ChatMessageResponse{}, nil
}
