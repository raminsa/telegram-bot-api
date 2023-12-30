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

	msg := tg.NewRevokeChatInviteLink()
	msg.Username = "username"
	msg.InviteLink = "url"

	_, err = tg.RevokeChatInviteLink(msg)
	if err != nil {
		log.Fatal(err)
	}
}
