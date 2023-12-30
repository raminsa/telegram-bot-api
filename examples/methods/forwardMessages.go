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

	msg := tg.NewForwardMessages()
	msg.ChatID = 1234
	msg.FromChatID = 1234
	msg.MessageIds = []int{1234, 1235}

	_, err = tg.ForwardMessages(msg)
	if err != nil {
		log.Fatal(err)
	}
}
