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

	message := tg.NewSetGameScore()
	message.UserID = 1234
	message.Score = 1
	message.InlineMessageID = "id"

	_, err = tg.SetGameScore(message)
	if err != nil {
		log.Fatal(err)
	}
}
