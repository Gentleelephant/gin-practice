package config

import (
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"io/ioutil"
)

var (
	DB           = &gorm.DB{}
	GlobalConfig = &Config{}
	ConfigPath   = "C:\\work\\code\\goPro\\gin-practice\\src\\config\\config.yaml"
)

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type LdapConf struct {
	Enabled              bool   `yaml:"enabled"`
	Url                  string `yaml:"url"`
	ReadTimeout          int    `yaml:"readTimeout"`
	StartTLS             bool   `yaml:"startTLS"`
	InsecureSkipVerify   bool   `yaml:"insecureSkipVerify"`
	RootCA               string `yaml:"rootCA"`
	RootCAData           string `yaml:"rootCAData"`
	ManagerDN            string `yaml:"managerDN"`
	ManagerPassword      string `yaml:"managerPassword"`
	UserSearchBase       string `yaml:"userSearchBase"`
	UserSearchFilter     string `yaml:"userSearchFilter"`
	GroupSearchBase      string `yaml:"groupSearchBase"`
	GroupSearchFilter    string `yaml:"groupSearchFilter"`
	UserMemberAttribute  string `yaml:"userMemberAttribute"`
	GroupMemberAttribute string `yaml:"groupMemberAttribute"`
	LoginAttribute       string `yaml:"loginAttribute"`
	MailAttribute        string `yaml:"mailAttribute"`
}

type MysqlConf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Config struct {
	Mysql  *MysqlConf `yaml:"mysql"`
	Server *Server    `yaml:"server"`
	LDAP   *LdapConf  `yaml:"ldap"`
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
