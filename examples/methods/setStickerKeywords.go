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

	message := tg.NewSetStickerKeywords()
	message.Sticker = "1234"
	message.Keywords = []string{"keyword 1", "keyword 2"}

	_, err = tg.SetStickerKeywords(message)
	if err != nil {
		log.Fatal(err)
	}
}
