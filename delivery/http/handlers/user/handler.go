package user_handler

import (
	userservice "hangout/service/user"
	"hangout/validator/uservalidator"
)

type Handler struct {
	validator uservalidator.Validator
	userSvc   userservice.Service
}

func New(v uservalidator.Validator, userSvc userservice.Service) Handler {
	return Handler{
		validator: v,
		userSvc:   userSvc,
	}
}
