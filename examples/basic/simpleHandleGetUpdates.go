package main

import (
	"fmt"
	"log"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	getUpdates := tg.NewGetUpdates()
	getUpdates.Offset = 1
	getUpdates.Timeout = 60
	getUpdates.Limit = 100

	updates := tg.GetUpdatesChan(getUpdates)
	for update := range updates {
		fmt.Println(update.UpdateID, getUpdates.Offset)
		if update.UpdateID >= getUpdates.Offset {
			getUpdates.Offset = update.UpdateID + 1
		}
		if update.Message != nil {
			fmt.Println(update.Message.Text)
		}
	}
}
