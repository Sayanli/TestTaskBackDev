package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP      `env-required:"true" yaml:"http"`
	MongoDB   `env-required:"true" yaml:"mongo_db"`
	JWTSecret `env-required:"true" yaml:"jwt_secret"`
}

type HTTP struct {
	Port string `env-required:"true" yaml:"port"`
}

type MongoDB struct {
	URL string `env-required:"true" yaml:"url"`
}

type JWTSecret struct {
	Secret string `env-required:"true" yaml:"secret"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig("config/config.yaml", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %w", err)
	}
	return &cfg, nil
}
