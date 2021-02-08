package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	IsDev      bool   `envconfig:"APP_MODE" default:"true"`
	GCPProject string `envconfig:"GCP_PROJECT" required:"true"`
	Slack
}

type Slack struct {
	ClientID        string `envconfig:"SLACK_CLIENT_ID" required:"true"`
	SecretID        string `envconfig:"SLACK_SECRET_ID" required:"true"`
	RedirectURL     string `envconfig:"SLACK_REDIRECT_URL" required:"true"`
	UserRedirectURL string `envconfig:"SLACK_USER_REDIRECT_URL" required:"true"`
}

func MustReadConfigFromEnv() *Config {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal(err)
		panic(err)
	}
	return &config
}
