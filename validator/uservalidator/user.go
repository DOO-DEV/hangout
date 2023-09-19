package uservalidator

import (
	"context"
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	param "hangout/param/http"
	customerr "hangout/pkg/error"
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

func (v Validator) checkUserExists(value interface{}) error {
	username := value.(string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	user, err := v.repo.GetUserByUsername(ctx, username)
	if errors.Is(err, customerr.RecordNotFoundErr) {
		return nil
	}

	if err == nil && user != nil {
		return customerr.UserExistErr
	}

	return fmt.Errorf("something went wrong")
}
