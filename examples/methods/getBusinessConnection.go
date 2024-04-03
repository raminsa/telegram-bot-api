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

	msg := tg.NewGetBusinessConnection()
	msg.BusinessConnectionId = "1234"

	_, err = tg.GetBusinessConnection(msg)
	if err != nil {
		log.Fatal(err)
	}
}
