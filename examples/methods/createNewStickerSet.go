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

	msg := tg.NewCreateNewStickerSet()
	msg.UserID = 1234
	msg.Name = "name"
	msg.Title = "title"
	sticker := types.InputSticker{
		Sticker:   tg.FileURL("url"),
		EmojiList: []string{"😀", "☺️"},
	}
	msg.Stickers = append(msg.Stickers, sticker)

	_, err = tg.CreateNewStickerSet(msg)
	if err != nil {
		log.Fatal(err)
	}
}
