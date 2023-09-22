package grouphandler

import (
	"github.com/labstack/echo/v4"
	authmiddleware "hangout/delivery/http/middlewares/auth"
)

func (h Handler) SetRoutes(g *echo.Group) {
	gWithAuth := g.Group("", authmiddleware.Auth(h.authSvc, h.authConfig))

	group := gWithAuth.Group("/groups")
	group.POST("", h.CreateGroup)
	group.GET("", h.ListAllGroups)
	group.GET("/my", h.GetMyGroup)

	request := gWithAuth.Group("/join_requests")
	request.POST("", h.JoinGroup)
	request.GET("", h.ListMyJoinRequest)
	request.GET("/group", h.ListJoinRequestToMyGroup)
	request.POST("/accept", h.AcceptJoin)

	conReq := gWithAuth.Group("/connection_requests")
	conReq.POST("", h.ConnectGroups)
	conReq.GET("", h.ListMyGroupConnections)
	conReq.POST("/accept", h.AcceptGroupConnection)
}
