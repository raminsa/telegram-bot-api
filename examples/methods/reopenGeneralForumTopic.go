package main

import (
	"log"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	msg := tg.NewReopenGeneralForumTopic()
	msg.Username = "username"

	_, err = tg.ReopenGeneralForumTopic(msg)
	if err != nil {
		log.Fatal(err)
	}
}
