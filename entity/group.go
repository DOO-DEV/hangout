package entity

type Group struct {
	ID        string
	CreatorID string
	Users     []User
}
