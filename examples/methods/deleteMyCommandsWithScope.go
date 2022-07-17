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

	deleteCommands := tg.NewDeleteMyCommandsWithScope(tg.NewBotCommandScopeDefault())

	_, err = tg.DeleteMyCommands(deleteCommands)
	if err != nil {
		log.Fatal(err)
	}
}
