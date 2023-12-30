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

	message := tg.NewDeleteMessages()
	message.ChatID = 1234
	message.MessageIds = []int{1234, 1235}

	_, err = tg.DeleteMessages(message)
	if err != nil {
		log.Fatal(err)
	}
}
