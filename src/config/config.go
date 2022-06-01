package config

import (
	"gin-practice/src/common"
	"gin-practice/src/config/ldap"
	"gin-practice/src/dao"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Mysql  *dao.MySQL `yaml:"mysql"`
	Server *Server    `yaml:"server"`
	LDAP   *ldap.LDAP `yaml:"ldap"`
}

// LoadConfig load config
func LoadConfig(path string) error {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, common.GlobalConfig)
	if err != nil {
		return err
	}
	return nil
}
