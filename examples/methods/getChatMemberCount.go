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

	msg := tg.NewGetChatMemberCount()
	msg.Username = "username"

	_, err = tg.GetChatMemberCount(msg)
	if err != nil {
		log.Fatal(err)
	}
}
