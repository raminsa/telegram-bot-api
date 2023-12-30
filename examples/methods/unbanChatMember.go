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

	msg := tg.NewUnbanChatMember()
	msg.Username = "username"
	msg.UserID = 1234

	_, err = tg.UnbanChatMember(msg)
	if err != nil {
		log.Fatal(err)
	}
}
