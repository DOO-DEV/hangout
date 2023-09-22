package chathandler

import (
	"github.com/labstack/echo/v4"
	authmiddleware "hangout/delivery/http/middlewares/auth"
)

func (h Handler) SetRoutes(g *echo.Group) {
	withAuth := g.Group("", authmiddleware.Auth(h.authSvc, h.authCfg))

	chat := withAuth.Group("/chats")
	chat.POST("/:id", h.ChatWithOtherUser)
	chat.GET("/:id", h.GetChatMessages)

}
