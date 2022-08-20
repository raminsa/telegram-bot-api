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

	message := tg.NewDeleteMessage()
	message.ChatID = 1234
	message.MessageID = 1234

	_, err = tg.DeleteMessage(message)
	if err != nil {
		log.Fatal(err)
	}
}
