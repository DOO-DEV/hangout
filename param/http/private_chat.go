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
