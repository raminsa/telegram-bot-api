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

	message := tg.NewCreateInvoiceLink()
	message.Title = "title"
	message.Description = "description"
	message.Payload = "payload"
	message.ProviderToken = "provider token"
	message.Currency = "currency"
	message.Prices = tg.NewLabeledPrices(
		tg.NewLabeledPrice("title1", 1234),
		tg.NewLabeledPrice("title2", 1234),
	)

	_, err = tg.CreateInvoiceLink(message)
	if err != nil {
		log.Fatal(err)
	}
}
