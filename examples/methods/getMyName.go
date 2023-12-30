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

	getMyName := tg.NewGetMyName()
	getMyName.LanguageCode = "en"

	_, err = tg.GetMyName(getMyName)
	if err != nil {
		log.Fatal(err)
	}
}
