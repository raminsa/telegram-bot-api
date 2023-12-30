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

	message := tg.NewCloseForumTopic()
	message.ChatID = 1234
	message.MessageThreadId = 1742

	_, err = tg.CloseForumTopic(message)
	if err != nil {
		log.Fatal(err)
	}
}
