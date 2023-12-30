package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/Raminsa/Telegram_API/telegram"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	tg.SetSecretToken("BotSecretToken") // optional

	hook := tg.NewSetWebhook()

	u, err := url.Parse("BotHookUrl")
	if err != nil {
		log.Fatal(err)
	}
	hook.URL = u
	hook.Certificate = tg.FilePath("CertFile")
	hook.MaxConnections = 100
	// hook.AllowedUpdates = []string{"message", "edited_channel_post", "callback_query"}

	webhook, err := tg.SetWebhook(hook)
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.MarshalIndent(webhook, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("status: \n%s\n", string(out))
}
