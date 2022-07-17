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

	getCommands := tg.NewGetMyCommandsWithScope(tg.NewBotCommandScopeDefault())

	_, err = tg.GetMyCommands(getCommands)
	if err != nil {
		log.Fatal(err)
	}
}
