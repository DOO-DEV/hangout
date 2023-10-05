package chatvalidator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	param "hangout/param/http"
)

func (v Validator) ValidatePrivateChatMessageRequest(req param.PrivateMessageRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Content),
		validation.Field(&req.Type),
		validation.Field(&req.Action, validation.Required),
		validation.Field(&req.ReceiverID, is.UUID),
	)
}
