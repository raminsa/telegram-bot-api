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

	msg := tg.NewSetChatPermissions()
	msg.Username = "username"

	permissions := tg.NewChatPermissions()
	permissions.CanChangeInfo = false
	permissions.CanSendMessages = true

	msg.Permissions = &permissions

	_, err = tg.SetChatPermissions(msg)
	if err != nil {
		log.Fatal(err)
	}
}
