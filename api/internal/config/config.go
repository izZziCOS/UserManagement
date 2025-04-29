package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	DBHost     string `env:"DB_HOST,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
	DBPort     string `env:"DB_PORT,required"`
	ServerPort string `env:"SERVER_PORT" envDefault:"8080"`
	Env        string `env:"ENV" envDefault:"development"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}