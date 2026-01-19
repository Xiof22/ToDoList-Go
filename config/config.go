package config

import (
	"github.com/Xiof22/ToDoList/internal/validator"
	"github.com/caarlos0/env"
)

type Config struct {
	Port             int    `env:"APP_PORT" envDefault:"8080"`
	TimezoneLocation string `env:"TIMEZONE_LOCATION" envDefault:"Asia/Ashgabat"`
	CookieStoreKey   string `env:"COOKIE_STORE_KEY,required" validate:"min=8"`
	SessionName      string `env:"SESSION_NAME,required"`
	AdminEmail       string `env:"ADMIN_EMAIL,required" validate:"email"`
	AdminPassword    string `env:"ADMIN_PASSWORD,required" validate:"min=4,max=8"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, validator.Validate.Struct(cfg)
}
