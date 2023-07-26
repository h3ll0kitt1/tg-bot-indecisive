package telegram

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(cmdList),
		tgbotapi.NewKeyboardButton(cmdSave),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(cmdDone),
		tgbotapi.NewKeyboardButton(cmdSurprise),
	),
)

func (bot *Bot) processCommand(m *tgbotapi.Message) error {

	switch command(m) {
	case cmdSave:
		bot.handleSetSave(m)
	case cmdDone:
		bot.handleSetDone(m)
	case cmdList:
		if err := bot.handleList(m); err != nil {
			return err
		}
	case cmdSurprise:
		if err := bot.handleSurprise(m); err != nil {
			return err
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
		bot.handleSave(m)

	case bot.Tracker.IsSet(m.Chat.ID, cmdDone) && m.Text != cmdDone:
		bot.handleDone(m)
	}

	bot.Tracker.UnSet(m.Chat.ID)

	msg := tgbotapi.NewMessage(m.Chat.ID, m.Text)
	msg.ReplyMarkup = numericKeyboard

	if _, err := bot.TgBot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (bot *Bot) handleSetSave(m *tgbotapi.Message) {
	m.Text = "какую книгу хотите добавить?"
	bot.Tracker.Update(m.Chat.ID, cmdSave)
}

func (bot *Bot) handleSetDone(m *tgbotapi.Message) {
	m.Text = "какую книгу вы прочитали?"
	bot.Tracker.Update(m.Chat.ID, cmdDone)
}

func (bot *Bot) handleList(m *tgbotapi.Message) error {
	ok, err := bot.Db.LenNotZero(m.Chat.ID)
	if err != nil {
		return err
	}

	if !ok {
		m.Text = "ваш список пуст"
		return nil
	}
	list, err := bot.Db.List(m.Chat.ID)
	m.Text = strings.Join(list, "\n")
	return nil
}

func (bot *Bot) handleSurprise(m *tgbotapi.Message) error {
	ok, err := bot.Db.LenNotZero(m.Chat.ID)
	if err != nil {
		return err
	}
	if !ok {
		m.Text = "ваш список пуст"
		return nil
	}
	surprise, err := bot.Db.Rand(m.Chat.ID)
	m.Text = surprise
	return nil
}

func (bot *Bot) handleSave(m *tgbotapi.Message) error {
	m.Text = unify(m.Text)
	ok, err := bot.Db.Save(m.Chat.ID, m.Text)
	if err != nil {
		return err
	}
	if ok {
		//m.Text = "книга " + `"` + m.Text + `"` + " добавлена"
		m.Text = "книга " + m.Text + " добавлена"
		return nil
	}
	m.Text = "эта книга и так уже у вас добалена..... мб прочитайте :/"
	return nil
}

func (bot *Bot) handleDone(m *tgbotapi.Message) error {
	m.Text = unify(m.Text)
	ok, err := bot.Db.Delete(m.Chat.ID, m.Text)
	if err != nil {
		return err
	}
	if ok {
		m.Text = "книга " + m.Text + " удалена"
		return nil
	}
	m.Text = "нельзя удалить чего нет"
	return nil
}
