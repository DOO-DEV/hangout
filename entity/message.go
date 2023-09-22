package entity

import "time"

type MsgType string

const (
	MsgTypeText = "text"
)

type Message struct {
	Sender    string
	Receiver  string
	Content   string
	Type      MsgType
	Timestamp time.Time
}
