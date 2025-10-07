package config

import (
	"fmt"
	"github.com/caarlos0/env"
        "github.com/joho/godotenv"
)

type Config struct {
	Port int `env:"APP_PORT" envDefault:"8080"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

        if err := godotenv.Load(); err != nil {
                fmt.Println("No .env file found, using system default envs")
        }

        if err := env.Parse(cfg); err != nil {
                return nil, err
        }

	return cfg, nil
}
