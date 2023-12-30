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

	message := tg.NewSendGame()
	message.ChatID = 1234
	message.GameShortName = "name"

	_, err = tg.SendGame(message)
	if err != nil {
		log.Fatal(err)
	}
}
