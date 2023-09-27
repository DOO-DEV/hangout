package main

import (
	minioadapter "hangout/adapter/minio"
	"hangout/config"
	"hangout/delivery/http"
	_ "hangout/docs"
	"hangout/repository/postgres"
	pgchat "hangout/repository/postgres/chat"
	pggroup "hangout/repository/postgres/group"
	pguser "hangout/repository/postgres/user"
	authservice "hangout/service/auth"
	chatservice "hangout/service/chat"
	groupservice "hangout/service/group"
	userservice "hangout/service/user"
	"hangout/validator/chatvalidator"
	"hangout/validator/groupvalidator"
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

// @title					Hangout
// @version				1.1
// @description			The HTTP documentation for Hangout API
// @termsOfService			http://swagger.io/terms/
// @license.name			Apache 2.0
// @schemes				http
// @host					localhost:3000
// @BasePath				/api/v1
// @securityDefinitions	bearerAuth
// @in						header
// @name					Authorization
// @description			Enter the token with the `Bearer ` prefix, e.g. `Bearer jwt_token_string`.
// @in						header
// @name					Authorization
// @description			Enter the token with the `Bearer ` prefix, e.g. `Bearer jwt_token_string`.
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
	imageStorage := minioadapter.New(cfg.Minio)
	userSvc := userservice.New(userRepo, authSvc, imageStorage)
	userValidator := uservalidator.New(userRepo)

	groupRepo := pggroup.New(pgDB)
	groupSvc := groupservice.New(groupRepo)
	groupValidator := groupvalidator.New()

	chatRepo := pgchat.New(pgDB)
	chatSvc := chatservice.New(chatRepo, groupRepo, userRepo)
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
