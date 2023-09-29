package messagehandler

import (
	authservice "hangout/service/auth"
	messageservice "hangout/service/message"
	"hangout/validator/messagevalidator"
)

type Handler struct {
	msgSvc       messageservice.Service
	authConfig   authservice.Config
	msgValidator messagevalidator.Validator
}

func New(msgSvc messageservice.Service, authCfg authservice.Config, msgValidator messagevalidator.Validator) Handler {
	return Handler{
		msgSvc:       msgSvc,
		authConfig:   authCfg,
		msgValidator: msgValidator,
	}
}
