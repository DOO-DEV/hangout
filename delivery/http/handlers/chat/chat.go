package chathandler

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	param "hangout/param/http"
	"hangout/pkg/claims"
	"hangout/pkg/httperr"
	"net/http"
	"sync"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var connections = make(map[string]*websocket.Conn)

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

			switch req.Action {
			case param.ActionSendTextMessage:
				go h.ReceivePrivateMessage(c.Request().Context(), conn, req, claims.ID)
			default:
				conn.WriteMessage(websocket.TextMessage, []byte("wrong action type"))
			}
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

	_, err = h.msgService.SavePrivateMessage(ctx, p)
	if err != nil {
		_, msg := httperr.Error(err)
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
		return
	}

	conn.WriteMessage(websocket.TextMessage, []byte("saved message"))

	return
}
