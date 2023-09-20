package userservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/password"
	"hangout/pkg/richerror"
)

func (s Service) Login(ctx context.Context, req param.LoginRequest) (*param.LoginResponse, error) {
	const op = "UserService.Login"

	user, err := s.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	if err := password.DecodePassword(req.Password, user.Password); err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindInvalid).WithMessage("username or password is wrong")
	}
	accessToken, err := s.authGenerator.CreateToken(user)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.LoginResponse{
		Firstname: user.FirsName,
		LastName:  user.LastName,
		Username:  user.Username,
		Token:     accessToken,
	}, nil

}
