package connections

import (

	"github.com/pkg/errors"
	"github.com/jinzhu/configor"

)

type Config struct {
	Mysql         *MysqlConfig   `yaml:"mysql"`

}

type MysqlConfig struct {
	Host 		string `yaml:"host"`
	Login 		string `yaml:"login"`
	Password 	string `yaml:"password"`
	DataBase 	string `yaml:"database"`
	Port 		int `yaml:"port"`
}

func LoadConfig(path string) (*Config, error) {
	c := &Config{}
	err := configor.Load(c, path)
	if err != nil {
		return nil, errors.Wrap(err, `load`)
	}
	return c, nil
}
