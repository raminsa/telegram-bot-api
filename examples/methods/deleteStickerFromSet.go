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

	message := tg.NewDeleteStickerFromSet()
	message.Sticker = "sticker"

	_, err = tg.DeleteStickerFromSet(message)
	if err != nil {
		log.Fatal(err)
	}
}
