package claims

import (
	"github.com/labstack/echo/v4"
	authservice "hangout/service/auth"
)

func GetClaimsFromEchoContext(c echo.Context, cfg authservice.Config) *authservice.Claims {
	return c.Get(cfg.AuthMiddlewareContextKey).(*authservice.Claims)
}
