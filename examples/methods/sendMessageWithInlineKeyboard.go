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

	msg := tg.NewSendMessage()
	msg.ChatID = 1234
	msg.Text = "text"
	keyboard := tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardCallbackData("btn1", "data1"),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardCallbackData("btn2", "data2"),
			tg.NewInlineKeyboardCallbackData("btn3", "data3"),
		),
	)

	msg.ReplyMarkup = keyboard

	_, err = tg.SendMessage(msg)
	if err != nil {
		log.Fatal(err)
	}
}
