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

	msg := tg.NewPinChatMessage()
	msg.Username = "username"
	msg.MessageID = 1234

	_, err = tg.PinChatMessage(msg)
	if err != nil {
		log.Fatal(err)
	}
}
