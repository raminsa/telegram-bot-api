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

	message := tg.NewSendSticker()
	message.ChatID = 1234
	message.Sticker = tg.FileURL("url")

	_, err = tg.SendSticker(message)
	if err != nil {
		log.Fatal(err)
	}
}
