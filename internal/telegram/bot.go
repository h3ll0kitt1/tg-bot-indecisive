package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/h3ll0kitt1/tg-bot-indecisive-helper/internal/config"
	"github.com/h3ll0kitt1/tg-bot-indecisive-helper/internal/storage"
	"github.com/h3ll0kitt1/tg-bot-indecisive-helper/internal/tracker"
)

type Bot struct {
	Db       storage.Storage
	TgBot    *tgbotapi.BotAPI
	Commands map[string]string
	Tracker  *tracker.Tracker
}

func NewBot(cfg config.Config, db storage.Storage, tracker *tracker.Tracker) (*Bot, error) {

	tgBot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}
	commands := setUpCommands()

	return &Bot{
		Db:       db,
		TgBot:    tgBot,
		Commands: commands,
		Tracker:  tracker,
	}, nil
}

func (bot *Bot) Run() error {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.TgBot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		if bot.isCommand(update.Message) {
			if err := bot.processCommand(update.Message); err != nil {
				return err
			}
			continue
		}

		if err := bot.processMessage(update.Message); err != nil {
			return err
		}
	}
	return nil
}
