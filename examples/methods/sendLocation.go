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

	msg := tg.NewSendLocation()
	msg.ChatID = 1234
	msg.Latitude = 1234
	msg.Longitude = 1234

	_, err = tg.SendLocation(msg)
	if err != nil {
		log.Fatal(err)
	}
}
