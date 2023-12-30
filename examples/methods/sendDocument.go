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

	msg := tg.NewSendDocument()
	msg.ChatID = 1234
	msg.Document = tg.FilePath("filePath")

	_, err = tg.SendDocument(msg)
	if err != nil {
		log.Fatal(err)
	}
}
