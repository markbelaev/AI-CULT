package main

import (
	"AI-CULT/internal/config"
	"AI-CULT/internal/handlers"
	"log/slog"
	"os"
	"time"

	"gopkg.in/telebot.v4"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	cfg := config.Load()

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

	slog.Info("Bot started")
	bot.Start()
}
