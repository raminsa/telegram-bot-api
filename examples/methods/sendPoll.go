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

	msg := tg.NewSendPoll()
	msg.ChatID = 1234
	msg.Question = "question"
	msg.Options = []string{"option1", "option2"}

	_, err = tg.SendPoll(msg)
	if err != nil {
		log.Fatal(err)
	}
}
