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

	getDefaultAdministratorRights := tg.NewGetMyDefaultAdministratorRights()

	_, err = tg.GetMyDefaultAdministratorRights(getDefaultAdministratorRights)
	if err != nil {
		log.Fatal(err)
	}
}
