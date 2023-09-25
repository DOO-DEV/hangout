package userservice

import (
	"context"
	param "hangout/param/http"
	"mime/multipart"
)

func (s Service) SaveProfileImage(
	ctx context.Context,
	_ param.SaveProfileImageRequest,
	image *multipart.FileHeader,
	userID string,
) (*param.SaveProfileImageRequest, error) {

	return nil, nil
}

func (s Service) GetPrimaryProfileImage(ctx context.Context, req param.GetPrimaryProfileImageRequest, userID string) (*param.GetPrimaryProfileImageResponse, error) {
	return "", nil
}

func (s Service) GetAllProfileImages(ctx context.Context, _ param.GetAllProfileImagesRequest, userID string) (*param.GetAllProfileImagesResponse, error) {
	return &param.GetAllProfileImageResponse{Data: []string{}}, nils
}

func (s Service) DeleteProfileImage(ctx context.Context, req param.DeleteProfileImage, userID string) (*param.DeleteProfileImage, error) {
	return nil, nil
}

func (s Service) SetImageAsPrimary(ctx context.Context, req param.SetImageAsPrimaryRequest, userID string) (*param.SetImageAsPrimaryResponse, error) {
	return nil, nil
}
