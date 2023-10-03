package privatechathandler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	param "hangout/param/http"
	"hangout/pkg/claims"
	"log"
	"net/http"
	"sync"
)

type client struct {
	conn    *websocket.Conn
	handler Handler
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var look = sync.RWMutex{}
var connections = make(map[string]*websocket.Conn)

func (h Handler) PrivateChaWsHandler(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	claims := claims.GetClaimsFromEchoContext(c, h.authCfg)

	look.RLock()
	connections[claims.ID] = conn
	look.RUnlock()

	userClient := client{conn: conn, handler: h}

	go userClient.readPump()
	go userClient.writePump()

	return nil
}

func (c client) readPump() {
	defer c.conn.Close()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
		}

		// turn message to json.
		// check the action of message
		// validate the payload of message
		// validate receiver
		// send the event to private-chat-service
		var prMsg param.PrivateMessageRequest
		if err := json.Unmarshal(message, prMsg); err != nil {
			return
		}

		if err := c.handler.validator.ValidatePrivateChatMessageRequest(prMsg); err != nil {
			c.conn.WriteJSON(err)
			return
		}

		switch {
		case prMsg.Action == param.SendPrivateMessageAction:
			// set an timeout context
			if prMsg.ReceiverID == "" {
				c.conn.WriteMessage(websocket.TextMessage, []byte("receiver_id can't be empty"))
				return
			}
			if _, err := c.handler.userSvc.GetUserByID(context.Background(), param.GetUserByIDRequest{UserID: prMsg.ReceiverID}); err != nil {
				c.conn.WriteMessage(websocket.TextMessage, []byte("receiver doesn't exist."))
				return
			}

			c.handler.chatSvc.CreatePrivateChat()
		}
	}
}
func (c client) writePump() {}
