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

	getMyShortDescription := tg.NewGetMyShortDescription()
	getMyShortDescription.LanguageCode = "en"

	_, err = tg.GetMyShortDescription(getMyShortDescription)
	if err != nil {
		log.Fatal(err)
	}
}
