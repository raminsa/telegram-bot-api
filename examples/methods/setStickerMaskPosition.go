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

	msg := tg.NewSetStickerMaskPosition()
	msg.Sticker = "1234"

	_, err = tg.SetStickerMaskPosition(msg)
	if err != nil {
		log.Fatal(err)
	}
}
