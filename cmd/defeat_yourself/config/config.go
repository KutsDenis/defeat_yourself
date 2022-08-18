package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Loader interface {
	Load() Config
}

type Config struct {
	Token string
	Admin int64
}

func (Config) Load() Config {
	path := "./configs/bot.json"

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
