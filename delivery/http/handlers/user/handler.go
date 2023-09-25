package user_handler

import (
	authservice "hangout/service/auth"
	userservice "hangout/service/user"
	"hangout/validator/uservalidator"
)

type Handler struct {
	validator uservalidator.Validator
	userSvc   userservice.Service
	authCfg   authservice.Config
}

func New(v uservalidator.Validator, userSvc userservice.Service, authCfg authservice.Config) Handler {
	return Handler{
		validator: v,
		userSvc:   userSvc,
		authCfg:   authCfg,
	}
}
