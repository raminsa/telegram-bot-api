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

	msg := tg.NewLeaveChat()
	msg.Username = "username"

	_, err = tg.LeaveChat(msg)
	if err != nil {
		log.Fatal(err)
	}
}
