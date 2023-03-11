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

	msg := tg.NewDeleteStickerSet()
	msg.Name = "name"

	_, err = tg.DeleteStickerSet(msg)
	if err != nil {
		log.Fatal(err)
	}
}
