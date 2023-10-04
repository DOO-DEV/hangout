package chatservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) UpsertPrivateChat(ctx context.Context, req param.UpsertPrivateChatRequest) (*param.UpsertPrivateChatResponse, error) {
	// check for chat existence
	// if not create a chat and update participants
	// call save message and save message
	// publish message to channel for subscribe

	const op = "ChatService.SendMessageToPrivateChat"

	chat, err := s.GetPrivateChatByName(ctx, param.GetPrivateChatByNameRequest{
		SenderID:   req.SenderID,
		ReceiverID: req.ReceiverID,
	})
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}
	if chat == nil {
		newChat, err := s.CreatePrivateChat(ctx, param.CreatePrivateChatRequest{
			Sender:   req.SenderID,
			Receiver: req.ReceiverID,
		})
		if err != nil {
			return nil, richerror.New(op).WithError(err)
		}
		// TODO - bring the db logic of this service to the create private chat.
		// TODO - do the above and blew function db operation in transaction
		_, err = s.InsertPrivateChatParticipants(ctx, param.InsertPrivateChatParticipantsRequest{
			ChatID:  newChat.ChatID,
			UserID1: req.SenderID,
			UserID2: req.ReceiverID,
		})
		if err != nil {
			return nil, richerror.New(op).WithError(err)
		}

		chat.ChatID = newChat.ChatID
		chat.ChatName = newChat.ChatName
	}

	return &param.UpsertPrivateChatResponse{
		ChatID:   chat.ChatID,
		ChatName: chat.ChatName,
	}, nil
}
