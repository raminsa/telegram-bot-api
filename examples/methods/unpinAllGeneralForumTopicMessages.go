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

	message := tg.NewUnpinAllGeneralForumTopicMessages()
	message.ChatID = 1234

	_, err = tg.UnpinAllGeneralForumTopicMessages(message)
	if err != nil {
		log.Fatal(err)
	}
}
