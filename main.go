package main

import (
	"hangout/config"
	"hangout/delivery/http"
)

func main() {
	cfg := config.Load()

	httpServer := http.New(cfg.HttpServer)

	httpServer.Serve()
}
