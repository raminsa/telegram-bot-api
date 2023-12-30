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

	setMyShortDescription := tg.NewSetMyShortDescription()
	setMyShortDescription.ShortDescription = "description"
	setMyShortDescription.LanguageCode = "en"

	_, err = tg.SetMyShortDescription(setMyShortDescription)
	if err != nil {
		log.Fatal(err)
	}
}
