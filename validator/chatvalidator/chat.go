package chatvalidator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	param "hangout/param/http"
)

func (v Validator) ValidateChatMessageRequest(req param.ChatMessageRequest) error {
	return validation.ValidateStruct(&req, validation.Field(&req.Content, validation.Required))
}
