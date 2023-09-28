package userservice

import (
	"context"
	param "hangout/param/http"
	"hangout/pkg/richerror"
	"mime/multipart"
	"sync"
)

func (s Service) SaveProfileImage(
	ctx context.Context,
	_ param.SaveProfileImageRequest,
	image *multipart.FileHeader,
	userID string,
) (*param.SaveProfileImageResponse, error) {
	const op = "UserService.SaveProfileImage"

	url, err := s.imageStorage.SaveImageIntoStorage(ctx, image)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	if err := s.repo.SaveProfileImageInfo(ctx, url, userID); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.SaveProfileImageResponse{ImageUrl: url}, nil
}

func (s Service) GetPrimaryProfileImage(ctx context.Context,
	_ param.GetPrimaryProfileImageRequest,
	userID string) (*param.GetPrimaryProfileImageResponse, error) {
	const op = "UserService.GetPrimaryProfileImage"

	fileName, err := s.repo.GetPrimaryProfileImage(ctx, userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	url, err := s.imageStorage.GetTemporaryProfileImageUrl(ctx, fileName)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}
	return &param.GetPrimaryProfileImageResponse{Url: url}, nil
}

func (s Service) GetAllProfileImages(ctx context.Context,
	_ param.GetAllProfileImagesRequest,
	userID string) (*param.GetAllProfileImagesResponse, error) {
	// TODO - consider to send back the more info about images. like created_at, id etc..
	// TODO - consider making an entity for profile images or not??
	const op = "UserService.GetAllProfileImages"

	names, err := s.repo.GetAllProfileImages(ctx, userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	var wg sync.WaitGroup
	urls := make([]string, len(names))
	for i, n := range names {
		wg.Add(1)
		go func(err error, i int, name string) {
			defer wg.Done()
			u, err := s.imageStorage.GetTemporaryProfileImageUrl(ctx, name)
			if err != nil {
				return
			}
			urls[i] = u
		}(err, i, n)
	}
	wg.Wait()

	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.GetAllProfileImagesResponse{Data: urls}, nil
}

func (s Service) DeleteProfileImage(ctx context.Context,
	req param.DeleteProfileImageRequest,
	userID string) (*param.DeleteProfileImageResponse, error) {
	const op = "UserService.DeleteProfileImage"

	name, err := s.repo.DeleteProfileImage(ctx, userID, req.ImageID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	if err := s.imageStorage.DeleteProfileImage(ctx, name); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.DeleteProfileImageResponse{}, nil
}

func (s Service) SetImageAsPrimary(ctx context.Context,
	req param.SetImageAsPrimaryRequest,
	userID string) (*param.SetImageAsPrimaryResponse, error) {
	const op = "UserService.SetImageAsPrimary"

	if err := s.repo.SetImageAsPrimary(ctx, userID, req.ImageID); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &param.SetImageAsPrimaryResponse{}, nil
}
