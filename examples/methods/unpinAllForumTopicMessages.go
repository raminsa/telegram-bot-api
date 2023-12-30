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

	msg := tg.NewUnpinAllForumTopicMessages()
	msg.Username = "username"
	msg.MessageThreadId = 1234

	_, err = tg.UnpinAllForumTopicMessages(msg)
	if err != nil {
		log.Fatal(err)
	}
}
