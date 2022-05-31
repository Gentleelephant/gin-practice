package config

import (
	"gin-practice/src/dao"
	yaml "gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"io/ioutil"
)

var (
	DB           *gorm.DB
	GolbalConfig = &Config{}
	ConfigPath   = "C:\\work\\code\\goPro\\gin-practice\\src\\config\\config.yaml"
)

type MySQL struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type LDAP struct {
	Enabled         bool   `yaml:"enabled"`
	Host            string `yaml:"url"`
	Port            string `yaml:"port"`
	ManagerDN       string `yaml:"managerDN"`
	ManagerPassword string `yaml:"managerPassword"`
	UserSearchBase  string `yaml:"userSearchBase"`
	LoginAttribute  string `yaml:"loginAttribute"`
	MailAttribute   string `yaml:"mailAttribute"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Mysql  *MySQL  `yaml:"mysql"`
	Server *Server `yaml:"server"`
	LDAP   *LDAP   `yaml:"ldap"`
}

// LoadConfig load config
func LoadConfig(path string) error {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, GolbalConfig)
	if err != nil {
		return err
	}

	// TODO 初始化gorm
	dao.InitDB()

	return nil
}
