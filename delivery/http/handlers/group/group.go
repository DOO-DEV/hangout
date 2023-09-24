package grouphandler

import (
	"github.com/labstack/echo/v4"
	param "hangout/param/http"
	"hangout/pkg/claims"
	"hangout/pkg/httperr"
	"net/http"
)

// CreateGroup godoc
//
//	@Summary		Create a group
//	@Description	Create a group for a user
//	@Security		auth
//	@Tags			group
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.CreateGroupRequest	true	"Create group"
//	@Success		201		{object}	param.CreteGroupResponse
//	@Router			/groups [post]
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

// ListAllGroups godoc
//
//	@Summary		Show all group
//	@Description	User can list all existence group
//	@Security		auth
//	@Tags			group
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.GetMyGroupRequest	true	"List Groups"
//	@Success		200		{object}	param.GetMyGroupResponse
//	@Router			/groups [get]
func (h Handler) ListAllGroups(c echo.Context) error {
	res, err := h.groupSvc.GetAllGroups(c.Request().Context(), param.GetAllGroupsRequest{})
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

// GetMyGroup godoc
//
//	@Summary		Get my group
//	@Description	Get created group info
//	@Security		auth
//	@Tags			group
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.GetMyGroupRequest	true	"My Group Info"
//	@Success		200		{object}	param.GetMyGroupResponse
//	@Router			/groups/my [get]
func (h Handler) GetMyGroup(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c, h.authConfig)
	res, err := h.groupSvc.GetMyGroup(c.Request().Context(), param.GetMyGroupRequest{}, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

// JoinGroup godoc
//
//	@Summary		Join group
//	@Description	User can join a group
//	@Security		auth
//	@Tags			group
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.JoinRequest	true	"Join to group"
//	@Success		2001	{object}	param.JoinResponse
//	@Router			/join_requests [post]
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

// ListMyJoinRequest godoc
//
//	@Summary		List join requests
//	@Description	Show list of all joining requests
//	@Security		auth
//	@Tags			group
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.ListJoinRequest	true	"List joins"
//	@Success		200		{object}	param.ListJoinRequestsResponse
//	@Router			/join_requests [get]
func (h Handler) ListMyJoinRequest(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c, h.authConfig)

	res, err := h.groupSvc.ListAllJoinRequests(c.Request().Context(), param.ListJoinRequest{}, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)

}

// ListJoinRequestToMyGroup godoc
//
//	@Summary		List admin join requests
//	@Description	Admin can see list of join requests to it's group
//	@Security		auth
//	@Tags			group
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.ListJoinRequestsToMyGroupRequest	true	"List to my group"
//	@Success		200		{object}	param.ListJoinRequestsToMyGroupResponse
//	@Router			/join_requests [get]
func (h Handler) ListJoinRequestToMyGroup(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c, h.authConfig)

	res, err := h.groupSvc.ListJoinRequestToMyGroup(c.Request().Context(), param.ListJoinRequestsToMyGroupRequest{}, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

// AcceptJoin godoc
//
//	@Summary		List admin join requests
//	@Description	Admin can see list of join requests to it's group
//	@Security		auth
//	@Tags			group
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.ListJoinRequestsToMyGroupRequest	true	"Accept"
//	@Success		200		{object}	param.ListJoinRequestsToMyGroupResponse
//	@Router			/join_requests [get]
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

// ConnectGroups godoc
//
//	@Summary		Connect groups
//	@Description	Admins can request to each other to connect their groups
//	@Security		auth
//	@Tags			group
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.GroupConnectionRequest	true	"Connect group"
//	@Success		201		{object}	param.GroupConnectionResponse
//	@Router			/connection_requests [post]
func (h Handler) ConnectGroups(c echo.Context) error {
	var req param.GroupConnectionRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.validator.ValidateGroupConnectionRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	claims := claims.GetClaimsFromEchoContext(c, h.authConfig)
	res, err := h.groupSvc.GroupConnectionRequest(c.Request().Context(), req, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusCreated, res)
}

// ListMyGroupConnections godoc
//
//	@Summary		List all group connections
//	@Description	Admins can see list of all other groups that connect with its group
//	@Security		auth
//	@Tags			group
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.MyGroupConnectionsRequest	true	"List connections"
//	@Success		201		{object}	param.MyGroupConnectionsResponse
//	@Router			/connection_requests [get]
func (h Handler) ListMyGroupConnections(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c, h.authConfig)

	res, err := h.groupSvc.ListMyGroupConnections(c.Request().Context(), param.MyGroupConnectionsRequest{}, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, res)
}

// AcceptGroupConnection godoc
//
//	@Summary		Accept group connection request
//	@Description	Admins can accept request from other group to join with its group
//	@Security		auth
//	@Tags			group
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.AcceptGroupConnectionRequest	true	"Accept group connections"
//	@Success		201		{object}	param.AcceptGroupConnectionResponse
//	@Router			/connection_requests/accept [get]
func (h Handler) AcceptGroupConnection(c echo.Context) error {
	var req param.AcceptGroupConnectionRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.validator.ValidateAcceptGroupConnection(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	claims := claims.GetClaimsFromEchoContext(c, h.authConfig)
	res, err := h.groupSvc.AcceptGroupConnectionToMyGroup(c.Request().Context(), req, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}
