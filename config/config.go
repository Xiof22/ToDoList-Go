package config

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"time"
)

type Config struct {
	Addr            string        `env:"APP_ADDR" envDefault:":8080"`
	ReadTimeout     time.Duration `env:"READ_TIMEOUT" envDefault:"5s"`
	WriteTimeout    time.Duration `env:"WRITE_TIMEOUT" envDefault:"10s"`
	IdleTimeout     time.Duration `env:"IDLE_TIMEOUT" envDefault:"120s"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"5s"`
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
