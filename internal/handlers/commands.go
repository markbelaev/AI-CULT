package handlers

import "gopkg.in/telebot.v4"

func handleStart(c telebot.Context) error {
	return c.Send("Привет, здесь собраны бесплатные AI-боты",
		AIKeyboard(),
	)
}

func handleAbout(c telebot.Context) error {
	return c.Send("Я разработчик")
}
