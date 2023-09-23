package chatservice

import (
	"context"
	"github.com/google/uuid"
	"hangout/entity"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) ChatWithOtherUser(ctx context.Context, req param.ChatMessageRequest, sender, receiver string) (*param.ChatMessageResponse, error) {
	const op = "ChatService.ChatWithOtherUser"

	// check the existence of receiver
	_, err := s.userRepo.GetUserByID(ctx, receiver)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	// check the chat existence
	chat, err := s.chatRepo.GetChatByUsersIds(ctx, sender, receiver)

	// if exist take chatID. if not create a chat and send back the ID
	if chat == nil {
		if err != nil {
			return nil, richerror.New(op).WithError(err)
		}
		chat, err = s.createChat(ctx, sender, receiver)
		if err != nil {
			return nil, richerror.New(op).WithError(err)
		}
	}

	m := entity.Message{
		ChatID:  chat.ID,
		Content: req.Content,
		Type:    entity.MsgTypeText,
	}
	// send message into messages table with chatID.
	if err := s.chatRepo.SaveMessage(ctx, m); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.ChatMessageResponse{}, nil
}

func (s Service) createChat(ctx context.Context, u1, u2 string) (*entity.Chat, error) {
	newChat := entity.Chat{
		ID:       uuid.NewString(),
		UsersIDs: []string{u1, u2},
		Type:     entity.ChatTypeNormal,
	}

	return s.chatRepo.CreateChat(ctx, &newChat)
}
