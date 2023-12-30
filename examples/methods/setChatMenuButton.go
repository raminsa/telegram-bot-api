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

	chatMenu := tg.NewSetChatMenuButton()
	chatMenu.MenuButton = tg.NewMenuButtonDefault()

	_, err = tg.SetChatMenuButton(chatMenu)
	if err != nil {
		log.Fatal(err)
	}
}
