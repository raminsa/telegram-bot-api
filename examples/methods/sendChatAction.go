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

	msg := tg.NewSendChatAction()
	msg.ChatID = 1234
	msg.Action = tg.TypingChatAction()

	_, err = tg.SendChatAction(msg)
	if err != nil {
		log.Fatal(err)
	}
}
