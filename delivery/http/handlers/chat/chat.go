package chathandler

import (
	"github.com/labstack/echo/v4"
	param "hangout/param/http"
	claims "hangout/pkg/claims"
	"hangout/pkg/httperr"
	"net/http"
)

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
