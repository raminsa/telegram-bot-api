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

	msg := tg.NewUnbanChatSenderChat()
	msg.Username = "username"
	msg.SenderChatID = 1234

	_, err = tg.UnbanChatSenderChat(msg)
	if err != nil {
		log.Fatal(err)
	}
}
