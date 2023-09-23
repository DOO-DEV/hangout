package user_handler

import (
	"github.com/labstack/echo/v4"
	param "hangout/param/http"
	"hangout/pkg/httperr"
	"net/http"
)

// Register godoc
//
//	@Summary		Register account
//	@Description	Create a new account for new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.RegisterRequest	true	"Create User"
//	@Success		201		{object}	param.RegisterResponse
//	@Router			/signup [post]
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
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
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
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}
