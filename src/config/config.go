package config

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"github.com/go-ldap/ldap/v3"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
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

func (l *LdapConf) NewLdapClient() (*ldap.Conn, error) {
	if l.InsecureSkipVerify {
		return ldap.DialURL(l.Url, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: l.InsecureSkipVerify}))
	}
	tlsConfig := tls.Config{}
	tlsConfig.RootCAs = x509.NewCertPool()
	var caCert []byte
	var err error
	// Load CA cert
	if l.RootCA != "" {
		if caCert, err = ioutil.ReadFile(l.RootCA); err != nil {
			log.Println("Failed to read CA cert", err)
			return nil, err
		}
	}
	if l.RootCAData != "" {
		if caCert, err = base64.StdEncoding.DecodeString(l.RootCAData); err != nil {
			log.Println("Failed to decode CA data", err)
			return nil, err
		}
	}
	if caCert != nil {
		tlsConfig.RootCAs.AppendCertsFromPEM(caCert)
	}
	conn, err := ldap.DialURL("tcp", ldap.DialWithTLSConfig(&tlsConfig))
	if err != nil {
		return nil, err
	}
	return conn, nil
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
