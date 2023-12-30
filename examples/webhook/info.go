package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	webhook, err := tg.GetWebhook()
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.MarshalIndent(webhook, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("webhook info: \n%s\n", string(out))
}
