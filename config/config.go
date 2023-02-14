package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func New() *Config {
	godotenv.Load()
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
