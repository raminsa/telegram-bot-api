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

	deleteCommands := tg.NewDeleteMyCommands()

	_, err = tg.DeleteMyCommands(deleteCommands)
	if err != nil {
		log.Fatal(err)
	}
}
