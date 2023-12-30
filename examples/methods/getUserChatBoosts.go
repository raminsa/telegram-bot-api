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

	msg := tg.NewGetUserChatBoosts()
	msg.ChatID = 1234
	msg.UserID = 1235

	_, err = tg.GetUserChatBoosts(msg)
	if err != nil {
		log.Fatal(err)
	}
}
