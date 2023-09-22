package groupservice

import (
	"context"
	"hangout/entity"
	dbparam "hangout/param/pgdb"
)

type Repository interface {
	CreateGroup(ctx context.Context, group entity.Group) error
	GetAllGroups(ctx context.Context) ([]*entity.Group, error)
	GetMyGroup(ctx context.Context, userID string) ([]entity.Member, error)
	CheckUserGroup(ctx context.Context, username string) (bool, error)
	AddToPendingList(ctx context.Context, list entity.PendingList) error
	ListJoinRequest(ctx context.Context, userID string) ([]entity.PendingList, error)
	GetOwnedGroup(ctx context.Context, userID string) (*entity.Group, error)
	ListAllJoinRequestToMyGroup(ctx context.Context, groupID string) ([]entity.PendingList, error)
	MoveFromPendingListToGroup(ctx context.Context, groupID string, userID string) error
	ConnectGroups(ctx context.Context, g1, g2 string) error
	ListMyGroupConnections(ctx context.Context, groupID string) ([]dbparam.GroupConnection, error)
	AcceptGroupConnection(ctx context.Context, from, to string) error
	CheckUserGroupConnection(ctx context.Context, u1, u2 string) (bool, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
