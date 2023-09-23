package entity

import "time"

type MsgType string

const (
	MsgTypeText = "text"
)

type Message struct {
	ID        string
	ChatID    string
	Content   string
	Type      MsgType
	Timestamp time.Time
}
