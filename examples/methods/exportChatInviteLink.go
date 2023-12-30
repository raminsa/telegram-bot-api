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

	msg := tg.NewExportChatInviteLink()
	msg.Username = "username"

	_, err = tg.ExportChatInviteLink(msg)
	if err != nil {
		log.Fatal(err)
	}
}
