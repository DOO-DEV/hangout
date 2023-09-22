package param

import (
	dbparam "hangout/param/pgdb"
	"time"
)

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

type JoinRequest struct {
	GroupID string `json:"group_id"`
}

type JoinResponse struct {
}

type ListJoinRequest struct {
}

type PendingJoinReq struct {
	SentAt time.Time `json:"sent_at"`
	Group  string    `json:"group"`
}
type ListJoinRequestsResponse struct {
	Data []PendingJoinReq `json:"data"`
}

type MemberRequestToMyGroup struct {
	User   string    `json:"user"`
	SentAt time.Time `json:"sent_at"`
}
type ListJoinRequestsToMyGroupRequest struct {
}

type ListJoinRequestsToMyGroupResponse struct {
	Data []MemberRequestToMyGroup `json:"data"`
}

type AcceptJoinRequest struct {
	UserID string `json:"user_id"`
}

type AcceptJoinResponse struct {
}

type GroupConnectionRequest struct {
	GroupID string `json:"group_id"`
}

type GroupConnectionResponse struct {
}

type MyGroupConnectionsRequest struct {
}

type MyGroupConnectionsResponse struct {
	Data []dbparam.GroupConnection `json:"data"`
}

type AcceptGroupConnectionRequest struct {
	GroupID string `json:"group_id"`
}

type AcceptGroupConnectionResponse struct {
}
