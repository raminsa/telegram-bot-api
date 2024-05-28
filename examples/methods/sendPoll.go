package main

import (
	"github.com/raminsa/telegram-bot-api/types"
	"log"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	msg := tg.NewSendPoll()
	msg.ChatID = 1234
	msg.Question = "question"

	var inputPollOptions []types.InputPollOption
	inputPollOptions = append(inputPollOptions, types.InputPollOption{
		Text: "text1",
	})
	inputPollOptions = append(inputPollOptions, types.InputPollOption{
		Text: "text2",
	})

	msg.Options = inputPollOptions

	_, err = tg.SendPoll(msg)
	if err != nil {
		log.Fatal(err)
	}
}
