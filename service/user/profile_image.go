package userservice

import (
	"context"
	param "hangout/param/http"
	"mime/multipart"
)

func (s Service) SaveProfileImage(ctx context.Context, image *multipart.FileHeader, userID string) error {

	return nil
}

func (s Service) GetPrimaryProfileImage(ctx context.Context, userID string) (string, error) {
	return "", nil
}

func (s Service) GetAllProfileImages(ctx context.Context, userID string) (*param.GetAllProfileImageResponse, error) {
	return &param.GetAllProfileImageResponse{Data: []string{}}, nils
}

func (s Service) DeleteProfileImage(ctx context.Context, req param.DeleteProfileImage, userID string) (*param.GetAllProfileImageResponse, error) {
	return nil, nil
}

func (s Service) SetImageAsPrimary(ctx context.Context, req param.SetImageAsPrimaryRequest, userID string) (*param.SetImageAsPrimaryResponse, error) {
	return nil, nil
}
