package param

import (
	"time"
)

type CreatePrivateChatRequest struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

type CreatePrivateChatResponse struct {
	ChatID   string `json:"chat_id"`
	ChatName string `json:"chat_name"`
}

type GetPrivateChatByNameRequest struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

type GetPrivateChatByNameResponse struct {
	ChatID   string
	ChatName string
}

type ChatMessageRequest struct {
	Content string `json:"content"`
}

type ChatMessageResponse struct {
}

type GetChatHistoryRequest struct {
}

type ChatMsg struct {
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type GetChatHistoryResponse struct {
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	Data     []ChatMsg `json:"data"`
}

type GetUserChatsRequest struct {
}

type GetUserChatResponse struct {
	Data []string `json:"data"`
}
