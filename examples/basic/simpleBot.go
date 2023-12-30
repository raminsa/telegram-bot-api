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

	me, err := tg.GetMe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("botID:", me.ID, "botUsername:", me.UserName)
}
