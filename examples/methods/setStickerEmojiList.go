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

	message := tg.NewSetStickerEmojiList()
	message.Sticker = "1234"
	message.EmojiList = []string{"😀", "☺️"}

	_, err = tg.SetStickerEmojiList(message)
	if err != nil {
		log.Fatal(err)
	}
}
