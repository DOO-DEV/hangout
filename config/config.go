package config

import "hangout/delivery/http"

type Config struct {
	Debug      bool        `koanf:"debug"`
	HttpServer http.Config `koanf:"http_server"`
}
