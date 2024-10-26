package config

import (
	"time"

	"github.com/caarlos0/env/v11"
	"go.uber.org/zap/zapcore"
)

type Logging struct {
	Level    zapcore.Level `env:"LEVEL" envDefault:"info"`
	Encoding string        `env:"ENCODING" envDefault:"json"`
}
type DB struct {
	DSN string `env:"DSN,required" envDefault:"postgres://shortener:shortener123@localhost:5432/shortener?sslmode=disable"`
}

type Server struct {
	Timeout time.Duration `env:"TIMEOUT" envDefault:"10s"`
	Port    string        `env:"PORT" envDefault:"8080"`
}

type Config struct {
	DB      DB      `envPrefix:"DB_"`
	Server  Server  `envPrefix:"SERVER_"`
	Logging Logging `envPrefix:"LOG_"`
}

func New() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
