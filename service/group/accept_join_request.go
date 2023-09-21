package groupservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) AcceptJoinRequest(ctx context.Context, req param.AcceptJoinRequest, adminID string) (*param.AcceptJoinResponse, error) {
	const op = "GroupService.AcceptJoinRequest"

	// check the user is admin
	// get the pending join requests from user to group(with groupID and userID)
	// delete this from pending list and change active status to false
	// join the user to given group
	gr, err := s.repo.GetOwnedGroup(ctx, adminID)
	if gr == nil {
		return nil, richerror.New(op).WithError(err)
	}

	if err := s.repo.MoveFromPendingListToGroup(ctx, gr.ID, req.UserID); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return nil, nil
}
