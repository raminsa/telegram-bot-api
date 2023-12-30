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

	msg := tg.NewGetUserChatBoosts()
	msg.ChatID = 1234
	msg.UserID = 1235

	_, err = tg.GetUserChatBoosts(msg)
	if err != nil {
		log.Fatal(err)
	}
}
