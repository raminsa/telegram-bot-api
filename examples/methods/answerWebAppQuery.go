package main

import (
	"log"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	message := tg.NewAnswerWebAppQuery()
	message.WebAppQueryID = "id"

	article := tg.NewInlineQueryResultArticle("1234", "title")
	article.Description = "description"
	article.InputMessageContent = tg.NewInputTextMessageContent()

	message.Result = article

	_, err = tg.AnswerWebAppQuery(message)
	if err != nil {
		log.Fatal(err)
	}
}
