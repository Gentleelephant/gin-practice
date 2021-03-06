package config

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"github.com/go-ldap/ldap/v3"
	"io/ioutil"
	"log"
)

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
