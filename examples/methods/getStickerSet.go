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

	message := tg.NewGetStickerSet()
	message.Name = "mySticker"

	_, err = tg.GetStickerSet(message)
	if err != nil {
		log.Fatal(err)
	}
}
