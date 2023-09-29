package entity

type GroupChatRole int

const (
	GroupChatOwnerRole GroupChatRole = iota + 1
	GroupChatAdminRole
	GroupChatMemberRole
)

type GroupChat struct {
	ID   string
	Name string
	Role GroupChatRole
}

func (c GroupChat) RoleIsValid() bool {
	if c.Role > GroupChatMemberRole || c.Role < GroupChatOwnerRole {
		return false
	}

	return true
}
