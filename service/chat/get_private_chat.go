package chatservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
	"slices"
	"strings"
)

func (s Service) GetPrivateChatByName(ctx context.Context, req param.GetPrivateChatByNameRequest) (*param.GetPrivateChatByNameResponse, error) {
	const op = "ChatService.GetPrivateChatByName"

	chatName := s.createPrivateChatName(req.Sender, req.Receiver)

	chat, err := s.chatRepo.GetPrivateChatByName(ctx, chatName)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}
	if chat == nil {
		return nil, nil
	}

	return &param.GetPrivateChatByNameResponse{
		ChatID:   chat.Name,
		ChatName: chat.ID,
	}, nil
}

func (s Service) createPrivateChatName(uid1, uid2 string) string {
	// the name of private chats are "userid-userid" and sort from higher-lower
	users := []string{uid1, uid2}
	slices.Sort(users)

	chatName := strings.Join(users, "-")

	return chatName
}
