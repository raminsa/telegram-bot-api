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

	msg := tg.NewSendMessage()
	msg.ChatID = 1234
	msg.ParseMode = tg.ModeMarkdown()
	msg.Text = tg.EscapeText(msg.ParseMode, "some text")

	_, err = tg.SendMessage(msg)
	if err != nil {
		log.Fatal(err)
	}
}
