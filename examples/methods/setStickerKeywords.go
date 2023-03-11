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

	msg := tg.NewSetStickerKeywords()
	msg.Sticker = "1234"
	msg.Keywords = []string{"ball"}

	_, err = tg.SetStickerKeywords(msg)
	if err != nil {
		log.Fatal(err)
	}
}
