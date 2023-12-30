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

	getMyDefaultAdministratorRights := tg.NewGetMyDefaultAdministratorRights()
	getMyDefaultAdministratorRights.ForChannels = true

	_, err = tg.GetMyDefaultAdministratorRights(getMyDefaultAdministratorRights)
	if err != nil {
		log.Fatal(err)
	}
}
