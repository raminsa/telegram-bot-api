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

	message := tg.NewEditGeneralForumTopic()
	message.ChatID = 1234
	message.Name = "name"

	_, err = tg.EditGeneralForumTopic(message)
	if err != nil {
		log.Fatal(err)
	}
}
