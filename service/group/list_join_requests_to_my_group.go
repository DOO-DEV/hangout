package groupservice

import (
	"context"
	"fmt"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) ListJoinRequestToMyGroup(ctx context.Context, _ param.ListJoinRequestsToMyGroupRequest, userID string) (*param.ListJoinRequestsToMyGroupResponse, error) {
	const op = "GroupService.ListJoinRequestToMyGroup"

	gr, err := s.repo.GetOwnedGroup(ctx, userID)
	if gr == nil {
		return nil, richerror.New(op).WithError(err)
	}
	fmt.Printf("%+v\n", gr)
	list, err := s.repo.ListAllJoinRequestToMyGroup(ctx, gr.ID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	pl := make([]param.MemberRequestToMyGroup, 0)
	for _, v := range list {
		p := param.MemberRequestToMyGroup{
			User:   v.UserID,
			SentAt: v.SentAt,
		}
		pl = append(pl, p)
	}

	res := &param.ListJoinRequestsToMyGroupResponse{Data: pl}

	return res, nil
}
