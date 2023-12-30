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

	msg := tg.NewGetCustomEmojiStickers()
	msg.CustomEmojiIds = []string{"😀", "☺️"}

	_, err = tg.GetCustomEmojiStickers(msg)
	if err != nil {
		log.Fatal(err)
	}
}
