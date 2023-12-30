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

	msg := tg.NewForwardMessage()
	msg.ChatID = 1234
	msg.FromChatID = 1234
	msg.MessageID = 1234

	_, err = tg.ForwardMessage(msg)
	if err != nil {
		log.Fatal(err)
	}
}
