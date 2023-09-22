package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	chathandler "hangout/delivery/http/handlers/chat"
	grouphandler "hangout/delivery/http/handlers/group"
	"hangout/delivery/http/handlers/health_check"
	user_handler "hangout/delivery/http/handlers/user"
	"hangout/pkg/constants"
	authservice "hangout/service/auth"
	chatservice "hangout/service/chat"
	groupservice "hangout/service/group"
	userservice "hangout/service/user"
	"hangout/validator/chatvalidator"
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
	chatHandler        chathandler.Handler
}

func New(httpCfg Config,
	userValidator uservalidator.Validator,
	userSvc userservice.Service,
	groupSvc groupservice.Service,
	gValidator groupvalidator.Validator,
	authSvc authservice.Service,
	authCfg authservice.Config,
	chatValidator chatvalidator.Validator,
	chatSvc chatservice.Service) Server {
	return Server{
		config:             httpCfg,
		router:             echo.New(),
		healthCheckHandler: health_check.New(),
		userHandler:        user_handler.New(userValidator, userSvc),
		groupHandler:       grouphandler.New(gValidator, groupSvc, authCfg, authSvc),
		chatHandler:        chathandler.New(chatValidator, authSvc, authCfg, chatSvc),
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
	s.chatHandler.SetRoutes(g)

	port := fmt.Sprintf(":%s", s.config.Port)
	s.router.Start(port)
}
