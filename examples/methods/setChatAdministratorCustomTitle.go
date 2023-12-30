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

	msg := tg.NewSetChatAdministratorCustomTitle()
	msg.Username = "username"
	msg.UserID = 1234
	msg.CustomTitle = "title"

	_, err = tg.SetChatAdministratorCustomTitle(msg)
	if err != nil {
		log.Fatal(err)
	}
}
