package param

import "time"

type GroupChatMessageRequest struct {
	ChatID   string `json:"chat_id"`
	SenderID string `json:"sender_id"`
	Content  string `json:"content"`
	Type     int    `json:"type"`
}

type GroupChatMessageResponse struct {
	Timestamp time.Time `json:"timestamp"`
	ID        string    `json:"id"`
}

type GetUnreadMessagesRequest struct {
	UserID string `json:"user_id"`
}

type GetUnreadMessagesResponse struct {
}

type SetStatusReadRequest struct {
	MessageID string `json:"message_id"`
}

type SetStatusReadResponse struct {
}
