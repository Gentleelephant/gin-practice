package config

import (
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type MySQL struct {
	Host string `yaml:"host"`

	Port string `yaml:"port"`

	User string `yaml:"user"`

	Password string `yaml:"password"`

	Database string `yaml:"database"`
}

type Config struct {
	Mysql *MySQL
}

// LoadConfig load config
func LoadConfig(path string) (*Config, error) {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("read config file error:", err)
		return nil, err
	}
	var conf Config
	err = yaml.Unmarshal(buf, &conf)

	if err != nil {
		log.Fatalln("unmarshal config file error:", err)
		return nil, err
	}

	return &conf, nil

}
