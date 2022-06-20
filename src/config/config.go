package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var (
	GlobalConfig = &Config{}
	ConfigPath   = "C:\\work\\code\\goPro\\gin-practice\\src\\config.yaml"
)

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Mysql  *MysqlConf   `yaml:"mysql"`
	Server *Server      `yaml:"server"`
	LDAP   *LdapConf    `yaml:"ldap"`
	Redis  *RedisConfig `yaml:"redis"`
}

// LoadConfig load config
func LoadConfig(path string) error {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, GlobalConfig)
	if err != nil {
		return err
	}
	return nil
}
