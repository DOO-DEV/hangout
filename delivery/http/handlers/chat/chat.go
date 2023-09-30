package chathandler

import (
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
				_, err := h.userSvc.GetUserByID(c.Request().Context(), param.GetUserByIDRequest{UserID: req.ReceiverID})
				if err != nil {
					_, msg := httperr.Error(err)
					conn.WriteMessage(websocket.TextMessage, []byte(msg))
					continue
				}

				chat, err := h.chatSvc.GetPrivateChatByName(c.Request().Context(), param.GetPrivateChatByNameRequest{
					SenderID:   claims.ID,
					ReceiverID: req.ReceiverID,
				})
				if err != nil {
					conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
					continue
				}
				if chat == nil {
					ch, err := h.chatSvc.CreatePrivateChat(c.Request().Context(), param.CreatePrivateChatRequest{
						Sender:   claims.ID,
						Receiver: req.ReceiverID,
					})
					if err != nil {
						conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
						continue
					}
					if _, err := h.chatSvc.InsertPrivateChatParticipants(c.Request().Context(), param.InsertPrivateChatParticipantsRequest{
						ChatID:  ch.ChatID,
						UserID1: claims.ID,
						UserID2: req.ReceiverID,
					}); err != nil {
						_, msg := httperr.Error(err)
						conn.WriteMessage(websocket.TextMessage, []byte(msg))
						continue
					}
					chat = &param.GetPrivateChatByNameResponse{
						ChatID:   ch.ChatID,
						ChatName: ch.ChatID,
					}
				}

				p := param.PrivateMessageRequest{
					ChatID:   chat.ChatID,
					SenderID: claims.ID,
					Content:  req.Content,
					Type:     req.Type,
				}

				_, err = h.msgService.SavePrivateMessage(c.Request().Context(), p)
				if err != nil {
					_, msg := httperr.Error(err)
					conn.WriteJSON(`{message:` + msg + `}`) // TODO - turn this to right json format
					continue
				}

				conn.WriteMessage(websocket.TextMessage, []byte("saved message"))

			//case param.ActionReadTextMessage:

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
