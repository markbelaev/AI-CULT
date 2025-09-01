package handlers

import "gopkg.in/telebot.v4"

func AIKeyboard() *telebot.ReplyMarkup {
	keyboard := &telebot.ReplyMarkup{
		ResizeKeyboard: true,
	}

	btnDeepSeek := keyboard.Text("DeepSeek")

	keyboard.Reply(
		keyboard.Row(btnDeepSeek),
	)

	return keyboard
}

func RemoveKeyboard() *telebot.ReplyMarkup {
	return &telebot.ReplyMarkup{
		RemoveKeyboard: true,
	}
}
