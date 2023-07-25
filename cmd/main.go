package main

import (
    "log"

    "github.com/h3ll0kitt1/tg-bot-indecisive-helper/internal/config"
    "github.com/h3ll0kitt1/tg-bot-indecisive-helper/internal/storage/redis"
    "github.com/h3ll0kitt1/tg-bot-indecisive-helper/internal/telegram"
    "github.com/h3ll0kitt1/tg-bot-indecisive-helper/internal/tracker"
)

func main() {

    cfg, err := config.New()
    if err != nil {
        log.Fatalf("Failed setup configuration: %v", err)
        return
    }

    db, err := redis.New(cfg)
    if err != nil {
        log.Fatalf("Failed to connect to Database: %v", err)
        return
    }

    tracker := tracker.New()

    bot, err := telegram.NewBot(cfg, db, tracker)
    if err != nil {
        log.Fatalf("Failed to connect to Telegram: %v", err)
        return
    }

    bot.TgBot.Debug = true

    if err := bot.Run(); err != nil {
        log.Fatalf("Failed to launch bot: %v", err)
        return
    }
}
