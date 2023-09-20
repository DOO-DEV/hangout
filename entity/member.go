package entity

import "time"

type Member struct {
	UserID   string
	GroupID  string
	JoinedAt time.Time
	Role     Role
}
