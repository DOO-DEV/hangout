package entity

import "time"

type User struct {
	ID        string
	Username  string
	FirsName  string
	LastName  string
	Password  string
	CreatedAt time.Time
}
