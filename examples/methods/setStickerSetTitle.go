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

	message := tg.NewSetStickerSetTitle()
	message.Name = "name"
	message.Title = "title"

	_, err = tg.SetStickerSetTitle(message)
	if err != nil {
		log.Fatal(err)
	}
}
