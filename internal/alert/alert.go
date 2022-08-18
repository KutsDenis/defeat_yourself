package alert

import (
	"defeat_yourself/cmd/defeat_yourself/config"
	"defeat_yourself/internal/messenger"
)

type Sender interface {
	ToAdmin(text string)
}

type Alert struct{}

func (Alert) ToAdmin(text string) {
	var msg messenger.Messenger
	var conf config.Loader

	msg = messenger.Message{}
	conf = config.Config{}

	send := msg.Send
	admin := conf.Load().Admin

	go send(admin, text)
}
