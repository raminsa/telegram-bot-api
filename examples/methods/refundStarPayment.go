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

	message := tg.NewRefundStarPayment()
	message.UserId = "user_id"
	message.TelegramPaymentChargeId = "id"

	_, err = tg.RefundStarPayment(message)
	if err != nil {
		log.Fatal(err)
	}
}
