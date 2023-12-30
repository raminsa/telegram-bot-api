package main

import (
	"fmt"
	"log"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	tg, err := telegram.NewWithBaseUrl("BotToken", "baseUrl")
	if err != nil {
		log.Fatal(err)
	}

	//active debug mode
	tg.Bot.Debug = true

	me, err := tg.GetMe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("botID:", me.ID, "botUsername:", me.UserName)

	//access debug log
	//method 1: write to console
	fmt.Println(tg.GetLoggerFile())

	//method 2: write to file
	err = tg.WriteLoggerFile("fileName")
	if err != nil {
		log.Fatal(err)
	}
}
