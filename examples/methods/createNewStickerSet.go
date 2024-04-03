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

	msg := tg.NewCreateNewStickerSet()
	msg.UserID = 1234
	msg.Name = "name"
	msg.Title = "title"

	stickers1 := tg.NewInputSticker()
	stickers1.EmojiList = []string{"🏀"}
	maskPosition := tg.NewMaskPosition()
	stickers1.MaskPosition = maskPosition
	stickers1.Keywords = []string{"ball"}
	stickers1.Sticker = tg.FileURL("url")

	stickers2 := tg.NewInputSticker()

	msg.Stickers = []any{stickers1, stickers2}

	_, err = tg.CreateNewStickerSet(msg)
	if err != nil {
		log.Fatal(err)
	}
}
