package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config - Estrutura que possui os dados de configuraçao
type Config struct {
	Server   string
	Port     string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("../config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
