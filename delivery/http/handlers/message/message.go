package messagehandler

import (
	"github.com/labstack/echo/v4"
	param "hangout/param/http"
	"hangout/pkg/claims"
	"hangout/pkg/httperr"
	"net/http"
)

func (h Handler) SendPrivateMessage(c echo.Context) error {
	var req param.PrivateMessageRequest

	claims := claims.GetClaimsFromEchoContext(c, h.authConfig)
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	req.SenderID = claims.ID
	if err := h.msgValidator.ValidateSendPrivateMessage(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := h.msgSvc.SavePrivateMessage(c.Request().Context(), req)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusCreated, res)
}
