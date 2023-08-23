package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

var Config *config

type config struct {
	HTTP      `env-required:"true" yaml:"http"`
	MongoDB   `env-required:"true" yaml:"mongo_db"`
	JWTSecret `env-required:"true" yaml:"jwt_secret"`
}

type HTTP struct {
	Port string `env-required:"true" yaml:"port"`
}

type MongoDB struct {
	URL      string `env-required:"true" yaml:"url"`
	Database string `env-required:"true" yaml:"database"`
}

type JWTSecret struct {
	Secret string `env-required:"true" yaml:"secret"`
}

func NewConfig() (*config, error) {
	var cfg config
	err := cleanenv.ReadConfig("config/config.yaml", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %w", err)
	}
	return &cfg, nil
}
