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

	message := tg.NewSetPassportDataErrors()
	message.UserID = 1234

	var Errors []interface{}

	passportElementErrorDataField := tg.PassportElementErrorDataField()
	passportElementErrorDataField.Source = "source"
	passportElementErrorDataField.Type = "type"
	passportElementErrorDataField.Message = "message"
	Errors = append(Errors, passportElementErrorDataField)

	message.Errors = Errors

	_, err = tg.SetPassportDataErrors(message)
	if err != nil {
		log.Fatal(err)
	}
}
