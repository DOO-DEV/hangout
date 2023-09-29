package chatservice

import (
	"context"
	param "hangout/param/http"
)

func (s Service) GetPrivateChatByName(ctx context.Context, req param.ChatMessageRequest)
