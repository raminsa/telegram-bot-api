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

	message := tg.NewHideGeneralForumTopic()
	message.ChatID = 1234

	_, err = tg.HideGeneralForumTopic(message)
	if err != nil {
		log.Fatal(err)
	}
}
