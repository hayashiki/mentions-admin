package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	IsDev bool `envconfig:"APP_MODE" default:"true"`
}

func MustReadConfigFromEnv() *Config {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal(err)
		panic(err)
	}
	return &config
}
