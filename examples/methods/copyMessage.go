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

	msg := tg.NewCopyMessage()
	msg.ChatID = 1234
	msg.FromChatID = 1234
	msg.MessageID = 1234

	_, err = tg.CopyMessage(msg)
	if err != nil {
		log.Fatal(err)
	}
}
