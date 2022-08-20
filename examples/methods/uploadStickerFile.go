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

	message := tg.NewUploadStickerFile()
	message.UserID = 1234
	message.PNGSticker = tg.FileURL("url")

	_, err = tg.UploadStickerFile(message)
	if err != nil {
		log.Fatal(err)
	}
}
