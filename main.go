package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6186964870:AAFsJIebRe3YrGKnV50bZybERt96YsOHI-A")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 120

	// updates := bot.GetUpdatesChan(u)

	msg := tgbotapi.NewMessage(6186964870, "<3")
	bot.Send(msg)
}
