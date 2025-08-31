package handlers

import (
	"log/slog"

	"gopkg.in/telebot.v4"
)

func RegisterAll(b *telebot.Bot) {
	registerCommands(b)
	slog.Info("All handlers registered")
}

func registerCommands(b *telebot.Bot) {
	b.Handle("/start", handleStart)
}
