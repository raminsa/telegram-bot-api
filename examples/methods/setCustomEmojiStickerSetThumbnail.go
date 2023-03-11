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

	msg := tg.NewSetCustomEmojiStickerSetThumbnail()
	msg.Name = "name"

	_, err = tg.SetCustomEmojiStickerSetThumbnail(msg)
	if err != nil {
		log.Fatal(err)
	}
}
