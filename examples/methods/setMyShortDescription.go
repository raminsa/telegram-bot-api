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

	setDescription := tg.NewSetMyShortDescription()
	setDescription.ShortDescription = "short description"

	_, err = tg.SetMyShortDescription(setDescription)
	if err != nil {
		log.Fatal(err)
	}
}
