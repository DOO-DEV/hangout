package auth

import (
	"github.com/google/uuid"
	"hangout/entity"
	param "hangout/param/http"
)

type Config struct {
}

type repository interface {
	Register(user *entity.User) error
}

type Service struct {
	repo repository
}

func New(repo repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Register(req param.RegisterRequest) (*param.RegisterResponse, error) {
	u := &entity.User{
		ID:       uuid.NewString(),
		FirsName: req.FirstName,
		LastName: req.LastName,
		Password: req.Password,
	}

	if err := s.repo.Register(u); err != nil {
		return nil, err
	}

	return &param.RegisterResponse{}, nil
}

func (s Service) Login() {

}
