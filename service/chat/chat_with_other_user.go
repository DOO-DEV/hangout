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
	// TODO - its better to have a chat table. first check the chat is exist and then save message to message table with their chatID

	// check the existence of sender and receiver
	// check connectivity of the users (for now. it will remove when policy change)
	// check the chat existence
	// if exist take chatID. if not create a chat and send back the ID
	// send message into messages table with chatID.

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
