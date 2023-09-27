package user_handler

import (
	"github.com/labstack/echo/v4"
	authmiddleware "hangout/delivery/http/middlewares/auth"
)

func (h Handler) SetRoutes(g *echo.Group) {
	g.POST("/signup", h.Register)
	g.POST("/login", h.Login)

	withAuth := g.Group("/user", authmiddleware.Auth(h.authSvc, h.authCfg))

	withAuth.POST("/profile_img", h.UploadProfileImage)
}
