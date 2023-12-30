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

	msg := tg.NewGetFile()
	msg.FileID = "file id"

	_, err = tg.GetFile(msg)
	if err != nil {
		log.Fatal(err)
	}
}
