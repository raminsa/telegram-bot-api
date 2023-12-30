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

	msg := tg.NewDeleteChatStickerSet()
	msg.Username = "username"

	_, err = tg.DeleteChatStickerSet(msg)
	if err != nil {
		log.Fatal(err)
	}
}
