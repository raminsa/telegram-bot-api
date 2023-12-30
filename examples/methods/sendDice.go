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

	msg := tg.NewSendDice()
	msg.ChatID = 1234
	msg.Emoji = "⚽"

	_, err = tg.SendDice(msg)
	if err != nil {
		log.Fatal(err)
	}
}
