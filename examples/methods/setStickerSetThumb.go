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

	message := tg.NewSetStickerSetThumb()
	message.UserID = 1234
	message.Name = "name"

	_, err = tg.SetStickerSetThumb(message)
	if err != nil {
		log.Fatal(err)
	}
}
