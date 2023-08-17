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

	message := tg.NewAnswerInlineQuery()
	message.InlineQueryID = "id"

	var ResultArticle []any

	article := tg.NewInlineQueryResultArticle("1234", "title")
	article.Description = "description"

	inputTextMessageContent := tg.NewInputTextMessageContent()
	inputTextMessageContent.Text = "text"

	article.InputMessageContent = inputTextMessageContent
	ResultArticle = append(ResultArticle, article)

	article = tg.NewInlineQueryResultArticle("1235", "title2")
	article.Description = "description"

	inputTextMessageContent = tg.NewInputTextMessageContent()
	inputTextMessageContent.Text = "text2"

	article.InputMessageContent = inputTextMessageContent
	ResultArticle = append(ResultArticle, article)

	message.Results = ResultArticle

	_, err = tg.AnswerInlineQuery(message)
	if err != nil {
		log.Fatal(err)
	}
}
