package groupvalidator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	param "hangout/param/http"
)

func (v Validator) ValidateJoinToGroupRequest(req param.JoinRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.GroupID, validation.Required),
	)
}

func (v Validator) ValidateAcceptJoinRequest(req param.AcceptJoinRequest) error {
	return validation.ValidateStruct(&req, validation.Field(&req.UserID, validation.Required))
}

func (v Validator) ValidateGroupConnectionRequest(req param.GroupConnectionRequest) error {
	return validation.ValidateStruct(&req, validation.Field(&req.GroupID, validation.Required))
}

func (v Validator) ValidateAcceptGroupConnection(req param.AcceptGroupConnectionRequest) error {
	return validation.ValidateStruct(&req, validation.Field(&req.GroupID, validation.Required))
}
