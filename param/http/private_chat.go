package param

import "time"

type PrivateMessageAction int

const (
	SendPrivateMessageAction PrivateMessageAction = iota + 1
	LeavePrivateMessageAction
)

type PrivateMessageRequest struct {
	Action     PrivateMessageAction `json:"action"`
	ReceiverID string               `json:"receiver_id"`
	Content    string               `json:"content"`
	Type       int                  `json:"type"`
}

type PrivateMessageResponse struct {
	ReceiverID string    `json:"receiver_id"`
	Timestamp  time.Time `json:"timestamp"`
	Content    string    `json:"content"`
	ID         string    `json:"id"`
}

type SavePrivateMessageRequest struct {
	SenderID string `json:"sender_id"`
	ChatID   string `json:"chat_id"`
	Content  string `json:"content"`
	Type     int    `json:"type"`
}

type SavePrivateMessageResponse struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

type UpsertPrivateChatRequest struct {
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"content"`
	Type       int    `json:"type"`
	SenderID   string `json:"sender_id"`
}

type UpsertPrivateChatResponse struct {
	ChatID   string `json:"chat_id"`
	ChatName string `json:"chat_name"`
}
