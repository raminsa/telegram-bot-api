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

	msg := tg.NewAnswerCallbackQuery()
	msg.CallbackQueryID = "id"

	_, err = tg.AnswerCallbackQuery(msg)
	if err != nil {
		log.Fatal(err)
	}
}
