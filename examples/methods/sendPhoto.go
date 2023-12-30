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

	msg := tg.NewSendPhoto()
	msg.ChatID = 1234
	msg.Photo = tg.FileURL("url")

	_, err = tg.SendPhoto(msg)
	if err != nil {
		log.Fatal(err)
	}
}
