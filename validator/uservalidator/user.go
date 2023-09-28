package uservalidator

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	param "hangout/param/http"
	"hangout/pkg/apperror"
	"time"
)

func (v Validator) ValidateRegisterRequest(req param.RegisterRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required),
		validation.Field(&req.Password, validation.Required),
		validation.Field(&req.FirstName, validation.Required),
		validation.Field(&req.LastName, validation.Required),
		validation.Field(&req.Username, validation.By(v.checkUserExists)),
	)
}

func (v Validator) ValidateLoginRequest(req param.LoginRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required),
		validation.Field(&req.Password, validation.Required),
	)
}

func (v Validator) ValidateDeleteProfileImageRequest(req param.DeleteProfileImageRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.ImageID, validation.Required, is.UUID),
	)
}

func (v Validator) ValidateSetImageAsPrimaryRequest(req param.SetImageAsPrimaryRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.ImageID, validation.Required),
	)
}

func (v Validator) checkUserExists(value interface{}) error {
	username := value.(string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	exists, err := v.repo.IsUserExists(ctx, username)
	if !exists && err == nil {
		return nil
	}

	if exists && err == nil {
		return apperror.UserExistErr
	}

	return fmt.Errorf("something went wrong")
}
