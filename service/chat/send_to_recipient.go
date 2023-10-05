package chatservice

import (
	"hangout/entity"
	param "hangout/param/http"
)

func (s Service) SendToRecipient(req param.SendToRecipientRequest) error {
	newMsg := entity.Message{
		ID:        req.ID,
		ChatID:    req.ChatID,
		SenderID:  req.SenderID,
		Content:   req.Content,
		Type:      entity.MsgType(req.Type),
		Status:    entity.MsgStatusDelivered,
		Timestamp: req.Timestamp,
	}
	if err := s.deliveryStorage.PublishToPrivateMessage(newMsg, req.ReceiverID); err != nil {
		return err
	}

	return nil
}

func (s Service) ListenForReceiveMessage(receiverID string) (*param.SendToRecipientResponse, error) {

	receiveChan := make(chan entity.Message)
	go s.deliveryStorage.SubscribeToPrivateMessage(receiverID, receiveChan)

	message := <-receiveChan

	return &param.SendToRecipientResponse{
		Content:   message.Content,
		Timestamp: message.Timestamp,
	}, nil
}
