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

	msg := tg.NewRestrictChatMember()
	msg.Username = "username"
	msg.UserID = 1234

	permissions := tg.NewChatPermissions()
	permissions.CanChangeInfo = false
	permissions.CanSendMessages = true

	msg.Permissions = permissions

	_, err = tg.RestrictChatMember(msg)
	if err != nil {
		log.Fatal(err)
	}
}
