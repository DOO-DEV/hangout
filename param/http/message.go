package param

import "time"

type PrivateMessageRequest struct {
	ChatID   string `json:"chat_id"`
	SenderID string `json:"sender_id"`
	Content  string `json:"content"`
	Type     int    `json:"type"`
}

type PrivateMessageResponse struct {
	Timestamp time.Time `json:"timestamp"`
	ID        string    `json:"id"`
}

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
