package main

import (
	"log"

	"github.com/Raminsa/Telegram_API/telegram"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	message := tg.NewEditMessageReplyMarkup()
	message.ChatID = 1234
	message.MessageID = 1234

	keyboard := tg.NewReplyKeyboardMarkup(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton("btn1"),
		),
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton("btn2"),
			tg.NewKeyboardButton("btn3"),
		),
	)
	keyboard.ResizeKeyboard = true
	keyboard.OneTimeKeyboard = true
	keyboard.Selective = true
	keyboard.InputFieldPlaceholder = "some text"

	message.ReplyMarkup = keyboard

	_, err = tg.EditMessageReplyMarkup(message)
	if err != nil {
		log.Fatal(err)
	}
}
