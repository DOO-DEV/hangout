package param

import "time"

type CreateGroupRequest struct {
	Name string `json:"name"`
}

type CreteGroupResponse struct {
	Name string `json:"name"`
}

type GetAllGroupsRequest struct {
}

type GroupInfo struct {
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllGroupsResponse struct {
	Data []GroupInfo `json:"data"`
}

type MemberInfo struct {
	User     string    `json:"user_id"`
	JoinedAt time.Time `json:"joined_at"`
	Role     string    `json:"role"`
}
type GetMyGroupRequest struct {
}

type GetMyGroupResponse struct {
	Group   string       `json:"group"`
	Members []MemberInfo `json:"members"`
}
