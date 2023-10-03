package chatservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) SendMessageToPrivateChat(ctx context.Context, req param.PrivateMessageRequest, senderID string) (*param.PrivateMessageResponse, error){
	// check for chat existence
	// if not create a chat and update participants
	// call save message and save message
	// publish message to channel for subscribe

	const op = "ChatService.SendMessageToPrivateChat"

	chat, err := s.GetPrivateChatByName(ctx, param.GetPrivateChatByNameRequest{
		SenderID:   senderID,
		ReceiverID: req.ReceiverID,
	})
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}
	if chat == nil {
		newChat , err := s.CreatePrivateChat(ctx, param.CreatePrivateChatRequest{
			Sender:   senderID,
			Receiver: req.ReceiverID,
		})
		if err != nil {
			return nil, richerror.New(op).WithError(err)
		}
		// TODO - bring the db logic of this service to the create private chat.
		// TODO - because it might this function break and participants never saved.
		// TODO - do the above and blew function db operation in transaction
		_, err = s.InsertPrivateChatParticipants(ctx, param.InsertPrivateChatParticipantsRequest{
			ChatID:  newChat.ChatID,
			UserID1: senderID,
			UserID2: req.ReceiverID,
		})
		if err != nil {
			return nil, richerror.New(op).WithError(err)
		}

		chat.ChatID = newChat.ChatID
		chat.ChatName = newChat.ChatName
	}

	if err := s.s(ctx, para)
}
