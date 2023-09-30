package chatvalidator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	param "hangout/param/http"
)

func (v Validator) ValidatePrivateChatMessageRequest(req param.PrivateChattingRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Content, validation.Required),
		validation.Field(&req.Type, validation.Required),
		validation.Field(&req.Action, validation.Required),
		validation.Field(&req.ReceiverID, validation.Required, is.UUID),
	)
}
