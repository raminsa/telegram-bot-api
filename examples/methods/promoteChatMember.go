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

	msg := tg.NewPromoteChatMember()
	msg.Username = "username"
	msg.UserID = 1234

	msg.CanEditMessages = false
	msg.CanInviteUsers = true

	_, err = tg.PromoteChatMember(msg)
	if err != nil {
		log.Fatal(err)
	}
}
