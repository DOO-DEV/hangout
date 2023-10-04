package redisadapter

import "fmt"

type Config struct {
	Port string `koanf:"port"`
	Addr string `koanf:"addr"`
}

type Adapter struct {
	rdb    *redis.Client
	config Config
}

func New(cfg Config) Adapter {
	client, err := redis.New()
	if err != nil {
		// TODO - decide to fatal or not?
		fmt.Printf("can't connect to redis: %s", err)
	}

	return Adapter{
		rdb:    client,
		config: cfg,
	}
}
