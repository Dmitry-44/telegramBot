package config

import (
	"github.com/Dmitry-44/telegramBot/connections"
	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
)

type Config struct {
	connections *connections.Config `yaml:"connections"`
}

func LoadConfig(path string) (*Config, error) {
	c := &Config{}
	err := configor.Load(c, path)
	if err != nil {
		return nil, errors.Wrap(err, `load`)
	}
	return c, nil
}

