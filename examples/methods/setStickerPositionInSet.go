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

	message := tg.NewSetStickerPositionInSet()
	message.Sticker = "🎯"
	message.Position = 0

	_, err = tg.SetStickerPositionInSet(message)
	if err != nil {
		log.Fatal(err)
	}
}
