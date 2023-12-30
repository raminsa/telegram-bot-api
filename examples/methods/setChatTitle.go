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

	msg := tg.NewSetChatTitle()
	msg.Username = "username"
	msg.Title = "title"

	_, err = tg.SetChatTitle(msg)
	if err != nil {
		log.Fatal(err)
	}
}
