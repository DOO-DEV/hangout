package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hangout/delivery/http/handlers/health_check"
	"hangout/pkg/constants"
)

type Config struct {
	Port string `json:"port"`
}

type Server struct {
	config             Config
	router             *echo.Echo
	healthCheckHandler health_check.Handler
}

func New(httpCfg Config) Server {
	return Server{
		config:             httpCfg,
		router:             echo.New(),
		healthCheckHandler: health_check.New(),
	}
}

func (s Server) Serve() {
	s.router.Use(middleware.RequestID())
	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())

	// Api prefix
	g := s.router.Group(constants.ApiEndpoint)

	s.healthCheckHandler.SetRoutes(g)

	port := fmt.Sprintf(":%s", s.config.Port)
	s.router.Start(port)
}
