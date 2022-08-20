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

	msg := tg.NewCreateChatInviteLink()
	msg.Username = "username"

	_, err = tg.CreateChatInviteLink(msg)
	if err != nil {
		log.Fatal(err)
	}
}
