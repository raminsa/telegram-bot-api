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

	setCommands := tg.NewSetMyCommandsWithScopeAndLanguage(tg.NewBotCommandScopeAllPrivateChats(), "en",
		tg.NewBotCommand("/command1", "description 1"),
		tg.NewBotCommand("/command2", "description 2"),
		tg.NewBotCommand("/command3", "description 3"),
	)

	_, err = tg.SetMyCommands(setCommands)
	if err != nil {
		log.Fatal(err)
	}
}
