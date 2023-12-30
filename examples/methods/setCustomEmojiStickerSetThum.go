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

	message := tg.NewSetCustomEmojiStickerSetThumb()
	message.Name = "name"
	message.CustomEmojiID = ""

	_, err = tg.SetCustomEmojiStickerSetThumb(message)
	if err != nil {
		log.Fatal(err)
	}
}
