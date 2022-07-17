package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Raminsa/Telegram_API/telegram"
)

func main() {
	fmt.Println("start at port:", "BotPortNumber")

	err := http.ListenAndServe("BotPortNumber", http.HandlerFunc(handleWebhook))
	if err != nil {
		log.Fatal(err)
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	update, err := telegram.HandleUpdate(r)
	if err != nil {
		telegram.HandleUpdateError(w, err)
		return
	}
	if update.Message != nil {
		fmt.Println(update.Message.Text)
	}
}
