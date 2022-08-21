package messenger

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Messenger interface {
	Send(Message)
}

type Message struct {
	ID   int64
	Text string
	Bot  *botapi.BotAPI
}

func (Message) Send(m Message) {
	msg := botapi.NewMessage(m.ID, m.Text)

	if _, err := m.Bot.Send(msg); err != nil {
		log.Fatal(err)
	}
}
