package commands

import (
	"defeat_yourself/internal/messenger"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander interface {
	CallCommand(command Command)
}

type Command struct {
	CMD string
	Usr *botapi.User
	Bot *botapi.BotAPI
}

func (Command) CallCommand(c Command) {
	var text string
	var msg messenger.Messenger
	msg = messenger.Message{}

	send := msg.Send

	switch c.CMD {
	case "help":
		text = "Что сюда написать?"
	case "start":
		text = "https://memepedia.ru/wp-content/uploads/2017/05/%D1%8F-%D1%81%D0%BA%D0%B0%D0%B7%D0%B0%D0%BB%D0%B0-%D1%81%D1%82%D0%B0%D1%80%D1%82%D1%83%D0%B5%D0%BC.jpg"
	default:
		text = "Неизвестная команда"
	}

	message := messenger.Message{
		ID:   c.Usr.ID,
		Text: text,
		Bot:  c.Bot,
	}

	send(message)
}
