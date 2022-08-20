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

	message := tg.NewEditMessageCaption()
	message.ChatID = 1234
	message.MessageID = 1742
	message.Caption = "text"

	_, err = tg.EditMessageCaption(message)
	if err != nil {
		log.Fatal(err)
	}
}
