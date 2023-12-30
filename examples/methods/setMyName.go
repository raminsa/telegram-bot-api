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

	setMyName := tg.NewSetMyName()
	setMyName.Name = "name"
	setMyName.LanguageCode = "en"

	_, err = tg.SetMyName(setMyName)
	if err != nil {
		log.Fatal(err)
	}
}
