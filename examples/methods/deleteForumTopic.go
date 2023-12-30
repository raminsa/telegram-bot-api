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

	msg := tg.NewDeleteForumTopic()
	msg.Username = "username"
	msg.MessageThreadID = 1234

	_, err = tg.DeleteForumTopic(msg)
	if err != nil {
		log.Fatal(err)
	}
}
