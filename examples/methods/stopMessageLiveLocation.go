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

	msg := tg.NewStopMessageLiveLocation()

	// use InlineMessageID
	msg.InlineMessageID = "id"

	// or use ChatID & MessageID
	msg.ChatID = 1234
	msg.MessageID = 1234

	_, err = tg.StopMessageLiveLocation(msg)
	if err != nil {
		log.Fatal(err)
	}
}
