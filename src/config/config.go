package config

import (
	yaml "gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

var GolbalConfig = &Config{}
var ConfigPath = "C:\\work\\code\\goPro\\gin-practice\\src\\config\\config.yaml"

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
	DB     *gorm.DB
}

func InitConfig() {
	config, err := LoadConfig(ConfigPath)
	if err != nil {
		log.Fatal("init config error:", err)
	}
	GolbalConfig = config
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
