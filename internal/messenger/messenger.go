package messenger

import (
	"defeat_yourself/cmd/defeat_yourself/config"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Messenger interface {
	Send(id int64, text string)
}

type Message struct {
}

func initBot() *botapi.BotAPI {
	var init config.Initializer

	init = config.Config{}
	bot := init.Init()

	return bot
}

func (m Message) Send(id int64, text string) {
	bot := initBot()
	msg := botapi.NewMessage(id, text)

	if _, err := bot.Send(msg); err != nil {
		log.Fatal(err)
	}
}
