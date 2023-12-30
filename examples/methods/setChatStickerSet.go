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

	msg := tg.NewSetChatStickerSet()
	msg.Username = "username"
	msg.StickerSetName = "name"

	_, err = tg.SetChatStickerSet(msg)
	if err != nil {
		log.Fatal(err)
	}
}
