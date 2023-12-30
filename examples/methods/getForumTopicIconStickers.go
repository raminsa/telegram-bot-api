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

	message := tg.NewGetForumTopicIconStickers()

	_, err = tg.GetForumTopicIconStickers(message)
	if err != nil {
		log.Fatal(err)
	}
}
