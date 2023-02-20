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

	msg := tg.NewUnpinAllForumTopicMessages()
	msg.Username = "username"
	msg.MessageThreadID = 1234

	_, err = tg.UnpinAllForumTopicMessages(msg)
	if err != nil {
		log.Fatal(err)
	}
}
