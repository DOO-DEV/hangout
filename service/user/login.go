package userservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/password"
)

func (s Service) Login(ctx context.Context, req param.LoginRequest) (*param.LoginResponse, error) {
	user, err := s.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	if err := password.DecodePassword(req.Password, user.Password); err != nil {
		return nil, err
	}
	accessToken, err := s.authGenerator.CreateToken(user)
	if err != nil {
		return nil, err
	}

	return &param.LoginResponse{
		Firstname: user.FirsName,
		LastName:  user.LastName,
		Username:  user.Username,
		Token:     accessToken,
	}, nil

}
