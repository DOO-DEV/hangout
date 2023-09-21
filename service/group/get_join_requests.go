package groupservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) ListAllJoinRequests(ctx context.Context, _ param.ListJoinRequest, userID string) (*param.ListJoinRequestsResponse, error) {
	const op = "GroupService.ListAllJoinRequests"

	list, err := s.repo.ListJoinRequest(ctx, userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	pd := make([]param.PendingJoinReq, 0)
	for _, v := range list {
		p := param.PendingJoinReq{
			SentAt: v.SentAt,
			Group:  v.GroupId,
		}
		pd = append(pd, p)
	}
	res := &param.ListJoinRequestsResponse{Data: pd}

	return res, nil
}
