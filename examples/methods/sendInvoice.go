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

	message := tg.NewSendInvoice()
	message.ChatIDStr = "chatId"
	message.Title = "title"
	message.Description = "description"
	message.Payload = "payload"
	message.Currency = "currency"
	message.Prices = tg.NewLabeledPrices(
		tg.NewLabeledPrice("title1", 1234),
		tg.NewLabeledPrice("title2", 1234),
	)

	_, err = tg.SendInvoice(message)
	if err != nil {
		log.Fatal(err)
	}
}
