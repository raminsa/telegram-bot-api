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

	msg := tg.NewGetFile()
	msg.FileID = "file id"

	_, err = tg.GetFile(msg)
	if err != nil {
		log.Fatal(err)
	}
}
