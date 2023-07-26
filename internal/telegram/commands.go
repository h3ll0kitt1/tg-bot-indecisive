package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	cmdSave     = "добавить"
	cmdDone     = "прочитано"
	cmdSurprise = "удиви меня"
	cmdList     = "мой список"
)

func (b *Bot) isCommand(m *tgbotapi.Message) bool {
	if _, ok := b.Commands[m.Text]; !ok {
		return false
	}
	return true
}

func command(m *tgbotapi.Message) string {
	return m.Text
}

func setUpCommands() map[string]string {
	cmd := map[string]string{
		cmdSave:     "Для сохранения книги в список нажмите " + cmdSave,
		cmdDone:     "Для удаления книги из списка нажмите " + cmdDone,
		cmdList:     "Для просмотра списка сохраненных книг нажмите " + cmdList,
		cmdSurprise: "Для выбора случайной книги нажмите " + cmdList,
	}
	return cmd
}
