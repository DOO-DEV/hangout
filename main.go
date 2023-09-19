package main

import (
	"hangout/config"
	"hangout/delivery/http"
	"hangout/repository/postgres"
	pguser "hangout/repository/postgres/user"
	authservice "hangout/service/auth"
	userservice "hangout/service/user"
	"hangout/validator/uservalidator"
)

type services struct {
	userValidator uservalidator.Validator
	userSvc       userservice.Service
}

func main() {
	cfg := config.Load()

	svc := setupServices(cfg)
	httpServer := http.New(cfg.HttpServer, svc.userValidator, svc.userSvc)

	httpServer.Serve()
}

func setupServices(cfg *config.Config) *services {
	pgDB := postgres.New(cfg.Postgres)

	userRepo := pguser.New(pgDB)
	authSvc := authservice.New(cfg.Auth)
	userSvc := userservice.New(userRepo, authSvc)
	userValidator := uservalidator.New(userRepo)

	return &services{
		userValidator: userValidator,
		userSvc:       userSvc,
	}
}
