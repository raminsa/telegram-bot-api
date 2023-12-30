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

	message := tg.NewUploadStickerFile()
	message.UserID = 1234
	message.Sticker = tg.FileURL("url")
	message.StickerFormat = "animated"

	_, err = tg.UploadStickerFile(message)
	if err != nil {
		log.Fatal(err)
	}
}
