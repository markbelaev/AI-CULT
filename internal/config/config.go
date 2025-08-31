package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TokenBot string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Error loading .env file", "err", err)
		os.Exit(1)
	}

	tokenBot := os.Getenv("TOKEN_BOT")
	if tokenBot == "" {
		slog.Error("TOKEN_BOT environment variable not set")
		os.Exit(1)
	}

	return &Config{
		TokenBot: tokenBot,
	}
}
