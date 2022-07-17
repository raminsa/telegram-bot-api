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

	msg := tg.NewSendAnimation()
	msg.ChatID = 1234
	msg.Animation = tg.FileURL("url")

	_, err = tg.SendAnimation(msg)
	if err != nil {
		log.Fatal(err)
	}
}
