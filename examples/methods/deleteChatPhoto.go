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

	msg := tg.NewDeleteChatPhoto()
	msg.Username = "username"

	_, err = tg.DeleteChatPhoto(msg)
	if err != nil {
		log.Fatal(err)
	}
}
