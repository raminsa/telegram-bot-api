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

	setDescription := tg.NewSetMyDescription()
	setDescription.Description = "description"

	_, err = tg.SetMyDescription(setDescription)
	if err != nil {
		log.Fatal(err)
	}
}
