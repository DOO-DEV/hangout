package chatservice

import "context"

func (s Service) GetUnreadPrivateChatMessages(ctx context.Context) {
	s.chatRepo.GetPrivateChatMessages()
}
