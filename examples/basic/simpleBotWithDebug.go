package main

import (
	"fmt"
	"log"

	"github.com/Raminsa/Telegram_API/telegram"
)

func main() {
	tg, err := telegram.NewWithBaseUrl("BotToken", "baseUrl")
	if err != nil {
		log.Fatal(err)
	}
	tg.Bot.Debug = true

	me, err := tg.GetMe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("botID:", me.ID, "botUsername:", me.UserName)

	err = tg.WriteLoggerFile("fileName")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tg.GetLoggerFile())
}
