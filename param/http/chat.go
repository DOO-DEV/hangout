package param

import (
	dbparam "hangout/param/pgdb"
	"time"
)

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
	Data []dbparam.Chat `json:"data"`
}
