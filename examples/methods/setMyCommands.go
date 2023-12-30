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

	setCommands := tg.NewSetMyCommandsScopeDefault()
	setCommands.LanguageCode = "en"

	_, err = tg.SetMyCommands(setCommands)
	if err != nil {
		log.Fatal(err)
	}
}
