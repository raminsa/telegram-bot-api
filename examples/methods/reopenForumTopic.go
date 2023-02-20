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

	msg := tg.NewReopenForumTopic()
	msg.Username = "username"
	msg.MessageThreadID = 1234

	_, err = tg.ReopenForumTopic(msg)
	if err != nil {
		log.Fatal(err)
	}
}
