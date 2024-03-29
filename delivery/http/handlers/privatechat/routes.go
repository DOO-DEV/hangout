package privatechathandler

import (
	"github.com/labstack/echo/v4"
	authmiddleware "hangout/delivery/http/middlewares/auth"
)

func (h Handler) SetRoutes(g *echo.Group) {
	g.GET("/chats", h.PrivateChaWsHandler, authmiddleware.Auth(h.authSvc, h.authCfg))
}
