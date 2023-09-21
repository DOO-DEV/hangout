package groupservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) GroupConnectionRequest(ctx context.Context, req param.GroupConnectionRequest, adminID string) (*param.GroupConnectionResponse, error) {
	const op = "GroupService.GroupConnectionRequest"

	// check it is admin -> get its group id
	// make db call to connect these group
	gr, err := s.repo.GetOwnedGroup(ctx, adminID)
	if gr == nil {
		return nil, richerror.New(op).WithError(err)
	}

	if err := s.repo.ConnectGroups(ctx, gr.ID, req.GroupID); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.GroupConnectionResponse{}, nil
}
