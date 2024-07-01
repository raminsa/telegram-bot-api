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

	msg := tg.NewSendPaidMedia()
	msg.ChatID = 1234
	msg.StarCount = 2

	var medias []any

	media := tg.NewInputPaidMediaPhoto()

	media.Media = tg.FileURL("url1")
	medias = append(medias, media)

	media.Media = tg.FileURL("url2")
	medias = append(medias, media)

	msg.Media = medias

	_, err = tg.SendPaidMedia(msg)
	if err != nil {
		log.Fatal(err)
	}
}
