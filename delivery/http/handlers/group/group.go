package grouphandler

import (
	"github.com/labstack/echo/v4"
	param "hangout/param/http"
	"hangout/pkg/claims"
	"hangout/pkg/httperr"
	"net/http"
)

func (h Handler) CreateGroup(c echo.Context) error {
	var req param.CreateGroupRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.validator.ValidateCreateGroupRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user := claims.GetClaimsFromEchoContext(c, h.authConfig)
	res, err := h.groupSvc.CreateGroup(c.Request().Context(), req, user.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h Handler) ListAllGroups(c echo.Context) error {
	res, err := h.groupSvc.GetAllGroups(c.Request().Context(), param.GetAllGroupsRequest{})
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

func (h Handler) GetMyGroup(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c, h.authConfig)
	res, err := h.groupSvc.GetMyGroup(c.Request().Context(), param.GetMyGroupRequest{}, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

func (h Handler) JoinGroup(c echo.Context) error {
	var req param.JoinRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.validator.ValidateJoinToGroupRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	claims := claims.GetClaimsFromEchoContext(c, h.authConfig)

	res, err := h.groupSvc.JoinGroup(c.Request().Context(), req, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h Handler) ListMyJoinRequest(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c, h.authConfig)

	res, err := h.groupSvc.ListAllJoinRequests(c.Request().Context(), param.ListJoinRequest{}, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)

}

func (h Handler) ListJoinRequestToMyGroup(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c, h.authConfig)

	res, err := h.groupSvc.ListJoinRequestToMyGroup(c.Request().Context(), param.ListJoinRequestsToMyGroupRequest{}, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

func (h Handler) AcceptJoin(c echo.Context) error {
	var req param.AcceptJoinRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.validator.ValidateAcceptJoinRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	claims := claims.GetClaimsFromEchoContext(c, h.authConfig)

	res, err := h.groupSvc.AcceptJoinRequest(c.Request().Context(), req, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}
