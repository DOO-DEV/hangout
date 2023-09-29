package entity

import "time"

type ChatType int

const (
	ChatTypePrivate ChatType = iota + 1
	ChatTypeGroup
)

type Chat struct {
	ID        string
	Type      ChatType
	Name      string
	CreatedAt time.Time
}

func (c Chat) TypeIsValid() bool {
	if c.Type > ChatTypeGroup || c.Type < ChatTypePrivate {
		return false
	}

	return true
}
