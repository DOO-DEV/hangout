package authmiddleware

import (
	"github.com/labstack/echo/v4"
	authservice "hangout/service/auth"
	"net/http"
	"strings"
)

func Auth(authSvc authservice.Service, cfg authservice.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			token := strings.TrimPrefix(authHeader, cfg.Prefix+" ")
			if len(token) < 1 {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			claims, err := authSvc.ParseToken(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			c.Set(cfg.AuthMiddlewareContextKey, claims)

			return next(c)
		}
	}
}
