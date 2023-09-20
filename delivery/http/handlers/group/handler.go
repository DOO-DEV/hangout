package grouphandler

import (
	authservice "hangout/service/auth"
	groupservice "hangout/service/group"
	"hangout/validator/groupvalidator"
)

type Handler struct {
	validator  groupvalidator.Validator
	groupSvc   groupservice.Service
	authConfig authservice.Config
	authSvc    authservice.Service
}

func New(v groupvalidator.Validator, svc groupservice.Service, authCfg authservice.Config, authSvc authservice.Service) Handler {
	return Handler{
		validator:  v,
		groupSvc:   svc,
		authConfig: authCfg,
		authSvc:    authSvc,
	}
}
