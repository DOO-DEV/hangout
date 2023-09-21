package entity

import "time"

type PendingList struct {
	UserID  string
	GroupId string
	SentAt  time.Time
	Active  bool
}
