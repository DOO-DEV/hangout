package user_handler

import (
	"github.com/labstack/echo/v4"
	param "hangout/param/http"
	"net/http"
)

func (h Handler) Register(c echo.Context) error {
	var req param.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.validator.ValidateRegisterRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := h.userSvc.Register(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h Handler) Login(c echo.Context) error {
	var req param.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.validator.ValidateLoginRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := h.userSvc.Login(c.Request().Context(), req)
	if err != nil {
		// TODO - find the under error and send right status and message
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
