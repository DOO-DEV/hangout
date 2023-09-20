package groupvalidator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	param "hangout/param/http"
)

func (v Validator) ValidateCreateGroupRequest(req param.CreateGroupRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required),
	)
}
