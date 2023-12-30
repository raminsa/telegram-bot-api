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

	msg := tg.NewSendPhoto()
	msg.ChatID = 1234
	msg.Photo = tg.FileURL("url")

	_, err = tg.SendPhoto(msg)
	if err != nil {
		log.Fatal(err)
	}
}
