package main

import (
	"log"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	msg := tg.NewSendMessage()
	msg.ChatID = 1234
	msg.Text = "text"

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

	msg.ReplyMarkup = keyboard

	_, err = tg.SendMessage(msg)
	if err != nil {
		log.Fatal(err)
	}
}
