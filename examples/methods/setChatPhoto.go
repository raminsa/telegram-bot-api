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

	msg := tg.NewSetChatPhoto()
	msg.Username = "username"
	msg.Photo = tg.FilePath("path")

	_, err = tg.SetChatPhoto(msg)
	if err != nil {
		log.Fatal(err)
	}
}
