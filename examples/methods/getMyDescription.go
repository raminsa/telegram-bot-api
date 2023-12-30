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

	getMyDescription := tg.NewGetMyDescription()
	getMyDescription.LanguageCode = "en"

	_, err = tg.GetMyDescription(getMyDescription)
	if err != nil {
		log.Fatal(err)
	}
}
