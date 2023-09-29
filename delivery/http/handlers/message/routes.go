package messagehandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(g *echo.Group) {
	msgGroup := g.Group("/messages")

	msgGroup.POST("", h.SendPrivateMessage)
}
