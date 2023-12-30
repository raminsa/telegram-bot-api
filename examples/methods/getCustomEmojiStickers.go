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

	message := tg.NewGetCustomEmojiStickers()

	message.CustomEmojiIds = []string{"1", "2", "3", "4"}

	_, err = tg.GetCustomEmojiStickers(message)
	if err != nil {
		log.Fatal(err)
	}
}
