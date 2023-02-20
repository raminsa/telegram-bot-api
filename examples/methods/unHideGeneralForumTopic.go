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

	msg := tg.NewUnHideGeneralForumTopic()
	msg.Username = "username"

	_, err = tg.UnHideGeneralForumTopic(msg)
	if err != nil {
		log.Fatal(err)
	}
}
