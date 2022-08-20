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

	msg := tg.NewSendVenue()
	msg.ChatID = 1234
	msg.Latitude = 1234
	msg.Longitude = 1234
	msg.Title = "title"
	msg.Address = "address"

	_, err = tg.SendVenue(msg)
	if err != nil {
		log.Fatal(err)
	}
}
