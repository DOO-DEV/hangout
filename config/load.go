package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
)

func Load() *Config {
	k := koanf.New(".")

	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Printf("error loading config.yml: %s", err)
	}

	cfg := &Config{}
	if err := k.Unmarshal("", cfg); err != nil {
		log.Fatalf("error unmarshaling config: %s", err)
	}

	if cfg.Debug {
		log.Printf("%+v", cfg)
	}

	return cfg
}
