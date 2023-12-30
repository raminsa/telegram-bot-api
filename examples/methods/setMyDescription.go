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

	setMyDescription := tg.NewSetMyDescription()
	setMyDescription.Description = "description"
	setMyDescription.LanguageCode = "en"

	_, err = tg.SetMyDescription(setMyDescription)
	if err != nil {
		log.Fatal(err)
	}
}
