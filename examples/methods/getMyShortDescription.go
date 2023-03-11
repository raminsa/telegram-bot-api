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

	getDescription := tg.NewGetMyShortDescription()

	_, err = tg.GetMyShortDescription(getDescription)
	if err != nil {
		log.Fatal(err)
	}
}
