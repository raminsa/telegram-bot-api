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

	msg := tg.NewSetChatDescription()
	msg.Username = "username"
	msg.Description = "description"

	_, err = tg.SetChatDescription(msg)
	if err != nil {
		log.Fatal(err)
	}
}
