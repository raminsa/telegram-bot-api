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

	message := tg.NewAddStickerToSet()
	message.UserID = 1234
	message.Name = "myName_by_<bot_username>"

	stickers := tg.NewInputSticker()
	stickers.EmojiList = []string{"🏀"}

	maskPosition := tg.NewMaskPosition()

	stickers.MaskPosition = maskPosition
	stickers.Keywords = []string{"ball"}
	stickers.Sticker = tg.FileURL("url")

	message.Sticker = stickers

	_, err = tg.AddStickerToSet(message)
	if err != nil {
		log.Fatal(err)
	}
}
