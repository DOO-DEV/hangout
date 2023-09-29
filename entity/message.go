package entity

import "time"

type MsgType int

const (
	MsgTypeText MsgType = iota + 1
	MsgTypePhoto
	MsgTypeVideo
	MsgTypeFile
)

type MsgStatus int

const (
	MsgStatusDelivered MsgStatus = iota + 1
	MsgStatusRead
)

type Message struct {
	ID        string
	ChatID    string
	SenderID  string
	Content   string
	Type      MsgType
	Status    MsgStatus
	Timestamp time.Time
}

func (m MsgType) TypeIsValid() bool {
	if m > MsgTypeFile || m < MsgTypeText {
		return false
	}

	return true
}
