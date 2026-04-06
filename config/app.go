package config

import "github.com/caarlos0/env/v11"

type AppConfig struct{}

func InitConfig() (AppConfig, error) {
	var cfg AppConfig

	err := env.Parse(&cfg)

	return cfg, err
}
