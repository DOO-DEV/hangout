package chatvalidator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	param "hangout/param/http"
)

func (v Validator) ValidatePrivateChatMessageRequest(req param.PrivateMessageRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Content, is.Alphanumeric),
		validation.Field(&req.Type, is.Int),
		validation.Field(&req.Action, validation.Required, is.Int),
		validation.Field(&req.ReceiverID, is.UUID),
	)
}
