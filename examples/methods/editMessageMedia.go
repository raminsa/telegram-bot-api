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

	message := tg.NewEditMessageMedia()
	message.ChatID = 1234
	message.MessageID = 1234

	media := tg.NewInputMediaPhoto()
	media.Media = tg.FileURL("url")
	message.Media = media

	_, err = tg.EditMessageMedia(message)
	if err != nil {
		log.Fatal(err)
	}
}
