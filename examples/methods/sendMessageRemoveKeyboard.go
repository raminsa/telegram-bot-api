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

	keyboard := tg.NewReplyKeyboardRemove(true)
	keyboard.Selective = true

	msg.ReplyMarkup = keyboard

	_, err = tg.SendMessage(msg)
	if err != nil {
		log.Fatal(err)
	}
}
