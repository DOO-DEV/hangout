package chathandler

import (
	authservice "hangout/service/auth"
	chatservice "hangout/service/chat"
	messageservice "hangout/service/message"
	userservice "hangout/service/user"
	"hangout/validator/chatvalidator"
)

type Handler struct {
	validator  chatvalidator.Validator
	authSvc    authservice.Service
	authCfg    authservice.Config
	chatSvc    chatservice.Service
	msgService messageservice.Service
	userSvc    userservice.Service
}

func New(v chatvalidator.Validator,
	authSvc authservice.Service,
	authCfg authservice.Config,
	chatSvc chatservice.Service,
	msgSvc messageservice.Service,
	userSvc userservice.Service,
) Handler {
	return Handler{
		validator:  v,
		authSvc:    authSvc,
		authCfg:    authCfg,
		chatSvc:    chatSvc,
		msgService: msgSvc,
		userSvc:    userSvc,
	}
}
