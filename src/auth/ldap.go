package auth

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"gin-practice/src/config"
	"gin-practice/src/global"
	"github.com/go-ldap/ldap/v3"
	"io/ioutil"
	"log"
)

type LdapProvider struct {
	Url                  string `yaml:"url"`
	ReadTimeout          int    `yaml:"readTimeout"`
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

func newProviderFromConf(c *config.LdapConf) *LdapProvider {
	return &LdapProvider{
		Url:                  c.Url,
		ReadTimeout:          c.ReadTimeout,
		InsecureSkipVerify:   c.InsecureSkipVerify,
		RootCA:               c.RootCA,
		RootCAData:           c.RootCAData,
		ManagerDN:            c.ManagerDN,
		ManagerPassword:      c.ManagerPassword,
		UserSearchBase:       c.UserSearchBase,
		UserSearchFilter:     c.UserSearchFilter,
		GroupSearchBase:      c.GroupSearchBase,
		GroupSearchFilter:    c.GroupSearchFilter,
		UserMemberAttribute:  c.UserMemberAttribute,
		GroupMemberAttribute: c.GroupMemberAttribute,
		LoginAttribute:       c.LoginAttribute,
		MailAttribute:        c.MailAttribute,
	}
}

func GetProvider() *LdapProvider {
	return newProviderFromConf(global.GlobalConfig.LDAP)
}

func (l *LdapProvider) newConn() (*ldap.Conn, error) {
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
	return ldap.DialURL("tcp", ldap.DialWithTLSConfig(&tlsConfig))
}

func (l *LdapProvider) Authentication(username, password string) (bool, string, error) {
	conn, err := l.newConn()
	if err != nil {
		log.Println("Failed to connect to LDAP server", err)
		return false, "", err
	}
	defer conn.Close()
	// Bind
	if err := conn.Bind(l.ManagerDN, l.ManagerPassword); err != nil {
		log.Println("Failed to bind to LDAP server", err)
		return false, "", err
	}
	// Search for user
	filter := fmt.Sprintf("(&(objectClass=organizationalPerson)(%s=%s))", l.LoginAttribute, username)
	log.Println("filter: ", filter)

	sql := ldap.NewSearchRequest(
		// ????????? basedn,????????????????????????????????????
		l.UserSearchBase,
		// ??????????????????????????? scope, derefAliases, sizeLimit, timeLimit,  typesOnly
		// ?????????????????? RFC4511 ????????????,???????????????
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		// ????????? LDAP ????????? Filter.??????????????????,?????????????????? uid=username ??? objectClass=organizationalPerson.
		// username ?????????????????????????????????
		//fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", "zhang"),
		filter,
		// ??????????????????????????????,?????????????????????.???????????????????????????????????????
		[]string{},
		nil,
	)

	sr, err := conn.Search(sql)
	if err != nil {
		log.Println("Query failed:", err)
		return false, "", err
	}

	dn := sr.Entries[0].DN
	fmt.Println("user dn: ", dn)

	err = conn.Bind(dn, password)
	if err != nil {
		log.Println("password error:", err)
		return false, "", err
	}
	return true, dn, nil
}
