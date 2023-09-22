package groupservice

import (
	"context"
	"errors"
	param "hangout/param/http"
	"hangout/pkg/errmsg"
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
	// can't connect with own groups
	if gr.ID == req.GroupID {
		wErr := richerror.New(op).WithError(errors.New("")).WithMessage(errmsg.ErrorMsgSelfGroupConnect)
		return nil, richerror.New(op).WithError(wErr).WithKind(richerror.KindInvalid)
	}

	if err := s.repo.ConnectGroups(ctx, gr.ID, req.GroupID); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.GroupConnectionResponse{}, nil
}
