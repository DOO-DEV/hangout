package groupservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) GetAllGroups(ctx context.Context, _ param.GetAllGroupsRequest) (*param.GetAllGroupsResponse, error) {
	const op = "GroupService.GetAllGroups"

	groups, err := s.repo.GetAllGroups(ctx)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	gr := make([]param.GroupInfo, 0)
	for _, val := range groups {
		g := param.GroupInfo{
			Name:      val.Name,
			Owner:     val.Owner,
			CreatedAt: val.CreatedAt,
		}

		gr = append(gr, g)
	}

	return &param.GetAllGroupsResponse{Data: gr}, nil
}
