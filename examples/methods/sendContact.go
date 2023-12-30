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

	msg := tg.NewSendContact()
	msg.ChatID = 1234
	msg.PhoneNumber = "phone number"
	msg.FirstName = "first name"

	_, err = tg.SendContact(msg)
	if err != nil {
		log.Fatal(err)
	}
}
