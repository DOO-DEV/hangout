package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hangout/delivery/http/handlers/health_check"
	user_handler "hangout/delivery/http/handlers/user"
	"hangout/pkg/constants"
	userservice "hangout/service/user"
	"hangout/validator/uservalidator"
)

type Config struct {
	Port string `koanf:"port"`
}

type Server struct {
	config             Config
	router             *echo.Echo
	healthCheckHandler health_check.Handler
	userHandler        user_handler.Handler
}

func New(httpCfg Config, userValidator uservalidator.Validator, userSvc userservice.Service) Server {
	return Server{
		config:             httpCfg,
		router:             echo.New(),
		healthCheckHandler: health_check.New(),
		userHandler:        user_handler.New(userValidator, userSvc),
	}
}

func (s Server) Serve() {
	s.router.Use(middleware.RequestID())
	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())

	// Api prefix
	g := s.router.Group(constants.ApiEndpoint)

	// Set up routes
	s.healthCheckHandler.SetRoutes(g)
	s.userHandler.SetRoutes(g)

	port := fmt.Sprintf(":%s", s.config.Port)
	s.router.Start(port)
}
