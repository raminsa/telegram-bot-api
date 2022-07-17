package main

import (
	"fmt"
	"log"

	"github.com/Raminsa/Telegram_API/telegram"
)

func main() {
	client := telegram.Client()

	client.BaseUrl = "baseUrl"
	client.Proxy = "proxy"
	client.ForceV4 = true
	client.DisableSSLVerify = true
	client.ForceAttemptHTTP2 = true

	tg, err := telegram.NewWithCustomClient("BotToken", &client)
	if err != nil {
		log.Fatal(err)
	}

	me, err := tg.GetMe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("botID:", me.ID, "botUsername:", me.UserName)
}
