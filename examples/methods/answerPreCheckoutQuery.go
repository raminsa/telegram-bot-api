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

	message := tg.NewAnswerShippingQuery()
	message.ShippingQueryID = "id"
	message.OK = false
	message.ErrorMessage = "err"

	_, err = tg.AnswerShippingQuery(message)
	if err != nil {
		log.Fatal(err)
	}
}
