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

	msg := tg.NewSendAudio()
	msg.ChatID = 1234
	msg.Audio = tg.FileURL("url")

	_, err = tg.SendAudio(msg)
	if err != nil {
		log.Fatal(err)
	}
}
