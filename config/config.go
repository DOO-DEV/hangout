package config

import (
	"hangout/delivery/http"
	"hangout/repository/postgres"
	authservice "hangout/service/auth"
)

type Config struct {
	Debug      bool               `koanf:"debug"`
	HttpServer http.Config        `koanf:"http_server"`
	Postgres   postgres.Config    `koanf:"postgres"`
	Auth       authservice.Config `koanf:"auth"`
}
