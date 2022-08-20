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

	msg := tg.NewEditMessageLiveLocation()

	// use InlineMessageID
	msg.InlineMessageID = "inline id"

	// or use ChatID & MessageID
	msg.ChatID = 1234
	msg.MessageID = 1234

	msg.Latitude = 1234
	msg.Longitude = 1234

	_, err = tg.EditMessageLiveLocation(msg)
	if err != nil {
		log.Fatal(err)
	}
}
