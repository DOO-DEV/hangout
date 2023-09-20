package groupservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) GetMyGroup(ctx context.Context, _ param.GetMyGroupRequest, owner string) (*param.GetMyGroupResponse, error) {
	const op = "GroupService.GetMyGroup"

	hasGroup, err := s.repo.CheckUserGroup(ctx, owner)
	if !hasGroup {
		return nil, richerror.New(op).WithError(err).WithMessage("you have not any group")
	}
	members, err := s.repo.GetMyGroup(ctx, owner)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	membersInfo := make([]param.MemberInfo, 0)
	for _, v := range members {
		m := param.MemberInfo{
			User:     v.UserID,
			JoinedAt: v.JoinedAt,
			Role:     string(v.Role),
		}
		membersInfo = append(membersInfo, m)
	}

	res := &param.GetMyGroupResponse{
		Group:   members[0].GroupID,
		Members: membersInfo,
	}

	return res, nil
}
