package grouphandler

import (
	"github.com/labstack/echo/v4"
	authmiddleware "hangout/delivery/http/middlewares/auth"
)

func (h Handler) SetRoutes(g *echo.Group) {
	group := g.Group("/groups", authmiddleware.Auth(h.authSvc, h.authConfig))

	group.POST("", h.CreateGroup)
	group.GET("", h.ListAllGroups)
	group.GET("/my", h.GetMyGroup)

	request := g.Group("/join_requests", authmiddleware.Auth(h.authSvc, h.authConfig))
	request.POST("", h.JoinGroup)
	request.GET("", h.ListMyJoinRequest)
	request.GET("/group", h.ListJoinRequestToMyGroup)
}
