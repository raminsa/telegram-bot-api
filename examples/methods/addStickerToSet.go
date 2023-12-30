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

	message := tg.NewAddStickerToSet()
	message.UserID = 1234
	message.Name = "myName_by_<bot_username>"
	message.Sticker = &types.InputSticker{
		Sticker:   tg.FileURL("url"),
		EmojiList: []string{"😀", "☺️"},
	}

	_, err = tg.AddStickerToSet(message)
	if err != nil {
		log.Fatal(err)
	}
}
