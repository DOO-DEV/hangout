package userservice

import (
	"context"
	"github.com/google/uuid"
	"hangout/entity"
	param "hangout/param/http"
	"hangout/pkg/password"
	"hangout/pkg/richerror"
)

func (s Service) Register(ctx context.Context, req param.RegisterRequest) (*param.RegisterResponse, error) {
	const op = "UserService.Register"

	hashedPassword, err := password.EncodePassword(req.Password)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}
	u := &entity.User{
		ID:       uuid.NewString(),
		FirsName: req.FirstName,
		LastName: req.LastName,
		Password: hashedPassword,
		Username: req.Username,
	}

	if err := s.repo.Register(ctx, u); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.RegisterResponse{}, nil
}
