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

	msg := tg.NewSendDice()
	msg.ChatID = 1234
	msg.Emoji = "⚽"

	_, err = tg.SendDice(msg)
	if err != nil {
		log.Fatal(err)
	}
}
