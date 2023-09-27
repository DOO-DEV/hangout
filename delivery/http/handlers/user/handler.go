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
	authSvc   authservice.Service
}

func New(v uservalidator.Validator, userSvc userservice.Service, authCfg authservice.Config, authSvc authservice.Service) Handler {
	return Handler{
		validator: v,
		userSvc:   userSvc,
		authCfg:   authCfg,
		authSvc:   authSvc,
	}
}
