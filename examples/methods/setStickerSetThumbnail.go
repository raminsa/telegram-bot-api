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

	message := tg.NewSetStickerSetThumbnail()
	message.UserID = 1234
	message.Name = "name"

	_, err = tg.SetStickerSetThumbnail(message)
	if err != nil {
		log.Fatal(err)
	}
}
