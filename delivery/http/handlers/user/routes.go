package user_handler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(g *echo.Group) {
	g.POST("/signup", h.Register)
	g.POST("/login", h.Login)
}
