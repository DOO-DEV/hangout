package entity

import "time"

type Group struct {
	ID        string
	Owner     string
	Name      string
	CreatedAt time.Time
}
