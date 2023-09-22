package groupservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) ListMyGroupConnections(ctx context.Context, _ param.MyGroupConnectionsRequest, adminID string) (*param.MyGroupConnectionsResponse, error) {
	const op = "GroupService.ListMyGroupConnections"

	gr, err := s.repo.GetOwnedGroup(ctx, adminID)
	if gr == nil {
		return nil, richerror.New(op).WithError(err)
	}
	list, err := s.repo.ListMyGroupConnections(ctx, gr.ID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.MyGroupConnectionsResponse{Data: list}, err
}
