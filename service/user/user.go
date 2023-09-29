package userservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
)

func (s Service) GetUserByID(ctx context.Context, req param.GetUserByIDRequest) (*param.GetUserByIDResponse, error) {
	const op = "UserService.GetUserByID"

	u, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.GetUserByIDResponse{
		UserID:    u.ID,
		Username:  u.Username,
		FirstName: u.FirsName,
		LastName:  u.LastName,
	}, nil
}
