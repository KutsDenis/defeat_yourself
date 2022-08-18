package config

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Initializer interface {
	Init() *botapi.BotAPI
}

func (c Config) Init() *botapi.BotAPI {
	var loader Loader

	loader = Config{}
	cfg := loader.Load()

	bot, err := botapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	return bot
}
