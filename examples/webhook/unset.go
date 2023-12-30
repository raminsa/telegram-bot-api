package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Raminsa/Telegram_API/telegram"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	hook := tg.NewDeleteWebhook()
	hook.DropPendingUpdates = true

	webhook, err := tg.DeleteWebhook(hook)
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.MarshalIndent(webhook, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("status: \n%s\n", string(out))
}
