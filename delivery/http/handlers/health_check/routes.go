package health_check

import (
	"github.com/labstack/echo/v4"
	"hangout/pkg/constants"
)

func (h Handler) SetRoutes(g *echo.Group) {
	g.GET(constants.HealthCheckEndpoint, h.HealthCheck)
}
