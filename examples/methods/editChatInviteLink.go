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

	msg := tg.NewEditChatInviteLink()
	msg.Username = "username"
	msg.InviteLink = "url"

	_, err = tg.EditChatInviteLink(msg)
	if err != nil {
		log.Fatal(err)
	}
}
