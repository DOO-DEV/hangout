package groupservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) AcceptGroupConnectionToMyGroup(ctx context.Context, req param.AcceptGroupConnectionRequest, adminID string) (*param.AcceptGroupConnectionResponse, error) {
	const op = "GroupService.AcceptGroupConnectionToMyGroup"

	// check group admin
	// check group connection request
	// turn accept to true
	gr, err := s.repo.GetOwnedGroup(ctx, adminID)
	if gr == nil {
		return nil, richerror.New(op).WithError(err)
	}

	if err := s.repo.AcceptGroupConnection(ctx, req.GroupID, gr.ID); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.AcceptGroupConnectionResponse{}, nil
}
