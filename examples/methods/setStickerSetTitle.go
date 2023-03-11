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

	msg := tg.NewSetStickerSetTitle()
	msg.Name = "name"
	msg.Title = "title"

	_, err = tg.SetStickerSetTitle(msg)
	if err != nil {
		log.Fatal(err)
	}
}
