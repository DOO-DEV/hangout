package redisadapter

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `koanf:"host"`
	DB       int    `koanf:"db"`
	Port     string `koanf:"port"`
	Addr     string `koanf:"addr"`
	Password string `koanf:"password"`
}

type Adapter struct {
	client *redis.Client
	config Config
}

func New(cfg Config) Adapter {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return Adapter{
		client: rdb,
	}
}

func (a Adapter) Client() *redis.Client {
	return a.client
}
