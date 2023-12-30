package main

import (
	"github.com/Raminsa/Telegram_API/types"
	"log"

	"github.com/Raminsa/Telegram_API/telegram"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	message := tg.NewSetStickerMaskPosition()
	message.Sticker = "1234"
	message.MaskPosition = types.MaskPosition{
		Point:  "forehead",
		XShift: -1.0,
		YShift: 1.0,
		Scale:  2.0,
	}

	_, err = tg.SetStickerMaskPosition(message)
	if err != nil {
		log.Fatal(err)
	}
}
