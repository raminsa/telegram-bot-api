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

	labeledPrices := tg.NewLabeledPrices()
	labeledPrice := tg.NewLabeledPrice()

	labeledPrice.Label = "label"
	labeledPrice.Amount = 123
	labeledPrices = append(labeledPrices, labeledPrice)

	message.Prices = labeledPrices

	_, err = tg.CreateInvoiceLink(message)
	if err != nil {
		log.Fatal(err)
	}
}
