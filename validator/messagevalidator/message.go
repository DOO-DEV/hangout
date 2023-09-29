package messagevalidator

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"hangout/entity"
	param "hangout/param/http"
)

func (v Validator) ValidateSendPrivateMessage(req param.PrivateMessageRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.SenderID, validation.Required, is.UUID),
		validation.Field(&req.Type, validation.Required, validation.By(v.isMessageTypeValid)),
		validation.Field(&req.Content, validation.Required),
		validation.Field(&req.ChatID, validation.Required, is.UUID),
	)
}

func (v Validator) isMessageTypeValid(value interface{}) error {
	msgType := value.(entity.MsgType)
	if msgType.TypeIsValid() {
		return nil
	}

	return fmt.Errorf("message type is not valid")
}
