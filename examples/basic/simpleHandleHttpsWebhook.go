package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	fmt.Println("start at port:", "BotPortNumber")

	err := http.ListenAndServeTLS("BotPortNumber", "BotCertFile", "BotKeyFile", http.HandlerFunc(handleWebhookTLS))
	if err != nil {
		log.Fatal(err)
	}
}

func handleWebhookTLS(w http.ResponseWriter, r *http.Request) {
	update, err := telegram.HandleUpdate(r)
	if err != nil {
		err = telegram.HandleUpdateError(w, err)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if update.Message != nil {
		fmt.Println(update.Message.Text)
	}
}
