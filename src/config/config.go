package config

import (
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
)

type MySQL struct {
	Host string `yaml:"host"`

	Port string `yaml:"port"`

	User string `yaml:"user"`

	Password string `yaml:"password"`

	Database string `yaml:"database"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Mysql  *MySQL
	Server *Server
}

// LoadConfig load config
func LoadConfig(path string) (*Config, error) {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = yaml.Unmarshal(buf, &conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil

}
