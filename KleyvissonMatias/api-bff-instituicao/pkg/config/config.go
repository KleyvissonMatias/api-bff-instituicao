package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   string
	Database string
}

func (config *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatal(err)
	}
}
