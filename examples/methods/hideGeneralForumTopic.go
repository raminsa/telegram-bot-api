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

	msg := tg.NewHideGeneralForumTopic()
	msg.Username = "username"

	_, err = tg.HideGeneralForumTopic(msg)
	if err != nil {
		log.Fatal(err)
	}
}
