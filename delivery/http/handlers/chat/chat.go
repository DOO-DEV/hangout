package chathandler

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	param "hangout/param/http"
	"hangout/pkg/claims"
	"hangout/pkg/httperr"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

type Message struct {
}

type Room struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
	Private    bool `json:"private"`
}
type Client struct {
	conn  *websocket.Conn
	send  chan []byte
	ID    string `json:"id"`
	Name  string `json:"name"`
	rooms map[*Room]bool
}

func newClient(conn *websocket.Conn, name, ID string) *Client {
	return &Client{
		conn:  conn,
		send:  make(chan []byte, 256),
		ID:    ID,
		Name:  name,
		rooms: make(map[*Room]bool),
	}
}

func (c *Client) readPump() {
	defer c.disconnect()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, jsonMsg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}
		c.handleNewMessage(jsonMsg)
	}
}
func (c *Client) writePump()                  {}
func (c *Client) disconnect()                 {}
func (c *Client) handleNewMessage(msg []byte) {}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var connections = make(map[string]*websocket.Conn)

var broadcast = make(chan param.PrivateMessageResponse)

func (h Handler) Chat(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c, h.authCfg)

	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer conn.Close()

	var connectionLocker sync.RWMutex
	connectionLocker.RLock()
	connections[claims.ID] = conn
	connectionLocker.RUnlock()

	go func() {
		for {
			var req param.PrivateChattingRequest
			if err := conn.ReadJSON(&req); err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
				continue
			}
			if err := h.validator.ValidatePrivateChatMessageRequest(req); err != nil {
				conn.WriteJSON(err)
				continue
			}

			go h.ReceivePrivateMessage(c.Request().Context(), conn, req, claims.ID)
			go h.sendPrivateMessage(c.Request().Context())
		}
	}()

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return nil
			}
		}
	}

	return nil
}

func (h Handler) ReceivePrivateMessage(ctx context.Context, conn *websocket.Conn, req param.PrivateChattingRequest, userID string) {
	_, err := h.userSvc.GetUserByID(ctx, param.GetUserByIDRequest{UserID: req.ReceiverID})
	if err != nil {
		_, msg := httperr.Error(err)
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
		return
	}

	chat, err := h.chatSvc.GetPrivateChatByName(ctx, param.GetPrivateChatByNameRequest{
		SenderID:   userID,
		ReceiverID: req.ReceiverID,
	})
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}
	if chat == nil {
		ch, err := h.chatSvc.CreatePrivateChat(ctx, param.CreatePrivateChatRequest{
			Sender:   userID,
			Receiver: req.ReceiverID,
		})
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			return
		}
		if _, err := h.chatSvc.InsertPrivateChatParticipants(ctx, param.InsertPrivateChatParticipantsRequest{
			ChatID:  ch.ChatID,
			UserID1: userID,
			UserID2: req.ReceiverID,
		}); err != nil {
			_, msg := httperr.Error(err)
			conn.WriteMessage(websocket.TextMessage, []byte(msg))
			return
		}
		chat = &param.GetPrivateChatByNameResponse{
			ChatID:   ch.ChatID,
			ChatName: ch.ChatID,
		}
	}

	p := param.PrivateMessageRequest{
		ChatID:   chat.ChatID,
		SenderID: userID,
		Content:  req.Content,
		Type:     req.Type,
	}

	msg, err := h.msgService.SavePrivateMessage(ctx, p)
	if err != nil {
		_, msg := httperr.Error(err)
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
		return
	}

	conn.WriteMessage(websocket.TextMessage, []byte("saved message"))

	broadcast <- param.PrivateMessageResponse{
		ReceiverID: req.ReceiverID,
		Timestamp:  msg.Timestamp,
		Content:    msg.Content,
		ID:         msg.ID,
	}

	return
}

func (h Handler) sendPrivateMessage(_ context.Context) {
	msg := <-broadcast
	conn, ok := connections[msg.ReceiverID]
	if !ok {
		return
	}

	if err := conn.WriteJSON(msg); err != nil {
		fmt.Println("cant send message to user: ", err)
		return
	}
}
