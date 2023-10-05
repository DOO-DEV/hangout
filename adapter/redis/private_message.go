package redisadapter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"hangout/entity"
	"time"
)

const (
	PrivateMessagesInbox = "private_messages"
)

func (a Adapter) PublishToPrivateMessage(message entity.Message, receiverID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	payload, _ := json.Marshal(message)
	topic := fmt.Sprintf("%s_%s", PrivateMessagesInbox, receiverID)
	if err := a.client.Publish(ctx, topic, payload).Err(); err != nil {
		log.Printf("publish error: %v\n", err)
		return err
	}

	return nil
}
func (a Adapter) SubscribeToPrivateMessage(userID string, msgChan chan<- entity.Message) {
	topic := fmt.Sprintf("%s_%s", PrivateMessagesInbox, userID)

	subscriber := a.client.Subscribe(context.Background(), topic)
	for {
		msg, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			fmt.Printf("subscribe error %v\n", err)
			continue
		}

		var message entity.Message
		json.Unmarshal([]byte(msg.Payload), &message)

		msgChan <- message
	}
}
