package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	grouphandler "hangout/delivery/http/handlers/group"
	"hangout/delivery/http/handlers/health_check"
	user_handler "hangout/delivery/http/handlers/user"
	"hangout/pkg/constants"
	authservice "hangout/service/auth"
	groupservice "hangout/service/group"
	userservice "hangout/service/user"
	"hangout/validator/groupvalidator"
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
	groupHandler       grouphandler.Handler
}

func New(httpCfg Config, userValidator uservalidator.Validator, userSvc userservice.Service, groupSvc groupservice.Service, gValidator groupvalidator.Validator, authSvc authservice.Service, authCfg authservice.Config) Server {
	return Server{
		config:             httpCfg,
		router:             echo.New(),
		healthCheckHandler: health_check.New(),
		userHandler:        user_handler.New(userValidator, userSvc),
		groupHandler:       grouphandler.New(gValidator, groupSvc, authCfg, authSvc),
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
	s.groupHandler.SetRoutes(g)

	port := fmt.Sprintf(":%s", s.config.Port)
	s.router.Start(port)
}
