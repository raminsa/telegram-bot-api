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

	msg := tg.NewSetStickerEmojiList()
	msg.Sticker = "1234"
	msg.EmojiList = []string{"🏀"}

	_, err = tg.SetStickerEmojiList(msg)
	if err != nil {
		log.Fatal(err)
	}
}
