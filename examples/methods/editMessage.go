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

	message := tg.NewEditMessageText()
	message.ChatID = 1234
	message.MessageID = 1234
	message.Text = "text"

	_, err = tg.EditMessageText(message)
	if err != nil {
		log.Fatal(err)
	}
}
