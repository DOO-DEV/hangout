package groupservice

import (
	"context"
	"github.com/google/uuid"
	"hangout/entity"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) CreateGroup(ctx context.Context, req param.CreateGroupRequest, owner string) (*param.CreteGroupResponse, error) {
	const op = "GroupService.CreateGroup"

	hasGroup, err := s.repo.CheckUserGroup(ctx, owner)
	if hasGroup {
		return nil, richerror.New(op).WithError(err).WithMessage("each user can join only one group")
	}
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}
	g := entity.Group{
		ID:    uuid.NewString(),
		Owner: owner,
		Name:  req.Name,
	}
	if err := s.repo.CreateGroup(ctx, g); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.CreteGroupResponse{
		Name: g.Name,
	}, nil
}
