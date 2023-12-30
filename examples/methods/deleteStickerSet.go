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

	message := tg.NewDeleteStickerSet()
	message.Name = "name"

	_, err = tg.DeleteStickerSet(message)
	if err != nil {
		log.Fatal(err)
	}
}
