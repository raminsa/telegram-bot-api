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

	message := tg.NewReplaceStickerInSet()
	message.UserID = 1234
	message.Name = "myName_by_<bot_username>"
	message.OldSticker = "1234"

	stickers := tg.NewInputSticker()
	stickers.EmojiList = []string{"🏀"}

	maskPosition := tg.NewMaskPosition()

	stickers.MaskPosition = maskPosition
	stickers.Keywords = []string{"ball"}
	stickers.Sticker = tg.FileURL("url")

	message.Sticker = stickers

	_, err = tg.ReplaceStickerInSet(message)
	if err != nil {
		log.Fatal(err)
	}
}
