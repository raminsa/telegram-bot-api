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

	message := tg.NewGetGameHighScores()
	message.UserID = 1234
	message.InlineMessageID = "id"

	_, err = tg.GetGameHighScores(message)
	if err != nil {
		log.Fatal(err)
	}
}
