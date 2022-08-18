package commands

import (
	"defeat_yourself/internal/messenger"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander interface {
	Command(cmd string, usr *botapi.User)
}

type Command struct {
}

func (c Command) Command(cmd string, usr *botapi.User) {
	var text string
	var msg messenger.Messenger
	msg = messenger.Message{}

	send := msg.Send

	switch cmd {
	case "help":
		text = "Что сюда написать?"
	case "start":
		text = "https://memepedia.ru/wp-content/uploads/2017/05/%D1%8F-%D1%81%D0%BA%D0%B0%D0%B7%D0%B0%D0%BB%D0%B0-%D1%81%D1%82%D0%B0%D1%80%D1%82%D1%83%D0%B5%D0%BC.jpg"
	default:
		text = "Неизвестная команда"
	}

	if text != "" {
		send(usr.ID, text)
	}
}
