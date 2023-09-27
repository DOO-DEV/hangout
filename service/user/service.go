package userservice

import (
	"context"
	"hangout/entity"
	"mime/multipart"
)

type Repository interface {
	Register(ctx context.Context, user *entity.User) error
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	SaveProfileImageInfo(ctx context.Context, imageUrl, useID string) error
}

type AuthGenerator interface {
	CreateToken(u *entity.User) (string, error)
}

type ImageStorage interface {
	SaveImageIntoStorage(ctx context.Context, file *multipart.FileHeader) (string, error)
}

type Service struct {
	repo          Repository
	authGenerator AuthGenerator
	imageStorage  ImageStorage
}

func New(repo Repository, authGen AuthGenerator, imageStorage ImageStorage) Service {
	return Service{
		repo:          repo,
		authGenerator: authGen,
		imageStorage:  imageStorage,
	}
}
