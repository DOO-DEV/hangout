package main

import (
	"hangout/config"
	"hangout/delivery/http"
	"hangout/repository/postgres"
	pgchat "hangout/repository/postgres/chat"
	pggroup "hangout/repository/postgres/group"
	pguser "hangout/repository/postgres/user"
	chatservice "hangout/service/chat"
	groupservice "hangout/service/group"
	"hangout/validator/chatvalidator"
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
	chatSvc        chatservice.Service
	chatValidator  chatvalidator.Validator
}

func main() {
	cfg := config.Load()

	svc := setupServices(cfg)
	httpServer := http.New(
		cfg.HttpServer,
		svc.userValidator,
		svc.userSvc,
		svc.groupSvc,
		svc.groupValidator,
		svc.authSvc,
		cfg.Auth,
		svc.chatValidator,
		svc.chatSvc,
	)

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

	chatRepo := pgchat.New(pgDB)
	chatSvc := chatservice.New(chatRepo, groupRepo)
	chatValidator := chatvalidator.New()
	return &services{
		userValidator:  userValidator,
		userSvc:        userSvc,
		groupSvc:       groupSvc,
		groupValidator: groupValidator,
		authSvc:        authSvc,
		chatSvc:        chatSvc,
		chatValidator:  chatValidator,
	}
}
