package main

import (
	"hangout/config"
	"hangout/delivery/http"
	"hangout/repository/postgres"
	pggroup "hangout/repository/postgres/group"
	pguser "hangout/repository/postgres/user"
	groupservice "hangout/service/group"
	"hangout/validator/groupvalidator"

	authservice "hangout/service/auth"
	userservice "hangout/service/user"
	"hangout/validator/uservalidator"
)

type services struct {
	userValidator  uservalidator.Validator
	userSvc        userservice.Service
	groupValidator groupvalidator.Validator
	groupSvc       groupservice.Service
	authSvc        authservice.Service
}

func main() {
	cfg := config.Load()

	svc := setupServices(cfg)
	httpServer := http.New(cfg.HttpServer, svc.userValidator, svc.userSvc, svc.groupSvc, svc.groupValidator, svc.authSvc, cfg.Auth)

	httpServer.Serve()
}

func setupServices(cfg *config.Config) *services {
	pgDB := postgres.New(cfg.Postgres)

	userRepo := pguser.New(pgDB)
	authSvc := authservice.New(cfg.Auth)
	userSvc := userservice.New(userRepo, authSvc)
	userValidator := uservalidator.New(userRepo)

	groupRepo := pggroup.New(pgDB)
	groupSvc := groupservice.New(groupRepo)
	groupValidator := groupvalidator.New()

	return &services{
		userValidator:  userValidator,
		userSvc:        userSvc,
		groupSvc:       groupSvc,
		groupValidator: groupValidator,
		authSvc:        authSvc,
	}
}
