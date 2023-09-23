package entity

import "time"

type ChatType string

const (
	ChatTypeNormal = "normal"
)

type Chat struct {
	ID        string
	UsersIDs  []string
	Type      ChatType
	CreatedAt time.Time
}
