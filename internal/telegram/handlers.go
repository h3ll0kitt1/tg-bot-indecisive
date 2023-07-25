package telegram

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	cmdSave     = "добавить"
	cmdDone     = "прочитано"
	cmdSurprise = "удиви меня"
	cmdList     = "мой список"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Мой список"),
		tgbotapi.NewKeyboardButton("Добавить"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Прочитано"),
		tgbotapi.NewKeyboardButton("Удиви меня"),
	),
)

func (bot *Bot) processCommand(m *tgbotapi.Message) error {

	switch m.Command() {

	case cmdSave:
		m.Text = "какую книгу хотите добавить?"
		bot.Tracker.Update(m.Chat.ID, cmdSave)

	case cmdDone:
		m.Text = "какую книгу вы прочитали?"
		bot.Tracker.Update(m.Chat.ID, cmdDone)

	case cmdList:
		list, err := bot.Db.List(m.Chat.ID)
		m.Text = strings.Join(list, "\n")
		if err == nil {
			m.Text = "ваш список пуст"
		}

	case cmdSurprise:
		surprise, err := bot.Db.Rand(m.Chat.ID)
		m.Text = surprise
		if err == nil {
			m.Text = "ваш список пуст"
		}
	}

	msg := tgbotapi.NewMessage(m.Chat.ID, m.Text)
	msg.ReplyMarkup = numericKeyboard

	if _, err := bot.TgBot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (bot *Bot) processMessage(m *tgbotapi.Message) error {

	switch {
	case bot.Tracker.NotSet(m.Chat.ID):
		m.Text = "кажется вам нужна помощь с ботом"

	case bot.Tracker.IsSet(m.Chat.ID, cmdSave) && m.Text != cmdSave:
		m.Text = "книга " + `"` + m.Text + `"` + " добавлена"

	case bot.Tracker.IsSet(m.Chat.ID, cmdDone) && m.Text != cmdDone:
		m.Text = "книга " + `"` + m.Text + `"` + " удалена"
	}

	bot.Tracker.UnSet(m.Chat.ID)

	msg := tgbotapi.NewMessage(m.Chat.ID, m.Text)
	msg.ReplyMarkup = numericKeyboard

	if _, err := bot.TgBot.Send(msg); err != nil {
		return err
	}
	return nil
}
