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
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
}

type GetPrivateChatByNameResponse struct {
	ChatID   string
	ChatName string
}

type InsertPrivateChatParticipantsRequest struct {
	ChatID  string `json:"chat_id"`
	UserID1 string `json:"user-id-1"`
	UserID2 string `json:"user_id_2"`
}

type InsertPrivateChatParticipantsResponse struct {
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

type PrivateChatAction string

const (
	ActionSendTextMessage PrivateChatAction = "send_txt_msg"
	ActionReadTextMessage PrivateChatAction = "read_txt_msg"
)

type PrivateChattingRequest struct {
	Action     PrivateChatAction `json:"action"`
	ChatID     string            `json:"chat_id"`
	ReceiverID string            `json:"receiver_id"`
	Content    string            `json:"content"`
	Type       int               `json:"type"`
}

type GetPrivateChatMessages struct {
	ChatID string `json:"chat_id"`
}

type PrivateChatMessages struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Content   string    `json:"content"`
	Status    int       `json:"status"`
}

type GetPrivateChatMessagesRepose struct {
	Data []PrivateChatMessages `json:"data"`
}

type GetUnreadPrivateChatMessagesRequest struct {
	ChatID string `json:"chat_id"`
}

type GetUnreadPrivateChatMessagesResponse struct {
	Data []PrivateChatMessages `json:"data"`
}
