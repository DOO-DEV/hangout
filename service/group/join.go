package groupservice

import (
	"context"
	"hangout/entity"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) JoinGroup(ctx context.Context, req param.JoinRequest, userID string) (*param.JoinResponse, error) {
	const op = "GroupService.JoinGroup"
	hasGroup, err := s.repo.CheckUserGroup(ctx, userID)
	if hasGroup {
		return nil, richerror.New(op).WithError(err).WithMessage("each user can join only one group")
	}

	p := entity.PendingList{
		UserID:  userID,
		GroupId: req.GroupID,
	}
	if err := s.repo.AddToPendingList(ctx, p); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.JoinResponse{}, nil
}
