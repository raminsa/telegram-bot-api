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

	getCommands := tg.NewGetMyCommandsScopeDefault()
	getCommands.LanguageCode = "en"

	_, err = tg.GetMyCommands(getCommands)
	if err != nil {
		log.Fatal(err)
	}
}
