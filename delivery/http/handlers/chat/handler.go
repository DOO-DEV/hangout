package chathandler

import (
	authservice "hangout/service/auth"
	chatservice "hangout/service/chat"
	"hangout/validator/chatvalidator"
)

type Handler struct {
	validator chatvalidator.Validator
	authSvc   authservice.Service
	authCfg   authservice.Config
	chatSvc   chatservice.Service
}

func New(v chatvalidator.Validator, authSvc authservice.Service, authCfg authservice.Config, chatSvc chatservice.Service) Handler {
	return Handler{validator: v, authSvc: authSvc, authCfg: authCfg, chatSvc: chatSvc}
}
