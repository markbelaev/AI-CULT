package main

import (
	"AI-CULT/internal/config"
	"AI-CULT/internal/handlers"
	"AI-CULT/internal/server"
	"log/slog"
	"os"
	"time"

	"gopkg.in/telebot.v4"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	cfg := config.Load()
	slog.Info("Config successfully loaded")

	s := server.New()

	go func() {
		if err := s.Start(); err != nil {
			slog.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	pref := telebot.Settings{
		Token: cfg.TokenBot,
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		slog.Error("Failed to create bot", "error", err)
		os.Exit(1)
	}

	handlers.RegisterAll(bot)
	slog.Info("All handlers registered")

	slog.Info("Bot started")
	bot.Start()
}
