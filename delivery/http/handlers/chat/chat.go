package chathandler

import (
	"github.com/labstack/echo/v4"
	param "hangout/param/http"
	claims "hangout/pkg/claims"
	"hangout/pkg/httperr"
	"net/http"
)

// ChatWithOtherUser godoc
//
//	@Summery		Chat with users
//	@Description	Chat with other users
//	@Security		auth
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"user id to chat with"
//	@Param			chat	body		param.ChatMessageRequest	true	"Chat message"
//	@Success		200		{object}	param.ChatMessageResponse
//	@Router			/chat [post]
func (h Handler) ChatWithOtherUser(c echo.Context) error {
	var req param.ChatMessageRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := h.validator.ValidateChatMessageRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	userIDToChatWith := c.Param("id")
	if userIDToChatWith == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	claims := claims.GetClaimsFromEchoContext(c, h.authCfg)

	res, err := h.chatSvc.ChatWithOtherUser(c.Request().Context(), req, claims.ID, userIDToChatWith)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

// GetChatMessages godoc
//
//	@Summery		Get chat history
//	@Description	History of chat
//	@Security		auth
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"user id to get chat history"
//	@Param			chat	body		param.GetChatHistoryRequest	false	"Chat message"
//	@Success		200		{object}	param.GetChatHistoryResponse
//	@Router			/chat [get]
func (h Handler) GetChatMessages(c echo.Context) error {
	userIDToChatWith := c.Param("id")
	if userIDToChatWith == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	claims := claims.GetClaimsFromEchoContext(c, h.authCfg)

	res, err := h.chatSvc.GetChatHistory(c.Request().Context(), param.GetChatHistoryRequest{}, claims.ID, userIDToChatWith)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

// GetUserChats godoc
//
//	@Summery		List chats
//	@Description	List all user chats
//	@Security		auth
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			chat	body		param.GetUserChatsRequest	false	"Chat message"
//	@Success		200		{object}	param.GetUserChatResponse
//	@Router			/chat [get]
func (h Handler) GetUserChats(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c, h.authCfg)

	res, err := h.chatSvc.ListUserChats(c.Request().Context(), param.GetUserChatsRequest{}, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}
