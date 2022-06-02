package auth

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
)

type CheckUser struct {
	Pool            *Pool
	ManagerDN       string
	ManagerPassword string
	UserSearchBase  string
	LoginAttribute  string
	MailAttribute   string
}

type LdapObject struct {
	DN       string
	Username string
	Email    string
}

func (u *CheckUser) Authentication(username string, password string) (*LdapObject, error) {

	conn, err := u.Pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	err = conn.Bind(u.ManagerDN, u.ManagerPassword)
	if err != nil {
		return nil, err
	}
	// Search for user
	filter := fmt.Sprintf("(&(objectClass=organizationalPerson)(%s=%s))", u.LoginAttribute, username)
	log.Println("filter: ", filter)

	sql := ldap.NewSearchRequest(
		// 这里是 basedn,我们将从这个节点开始搜索
		u.UserSearchBase,
		// 这里几个参数分别是 scope, derefAliases, sizeLimit, timeLimit,  typesOnly
		// 详情可以参考 RFC4511 中的定义,文末有链接
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		// 这里是 LDAP 查询的 Filter.这个例子例子,我们通过查询 uid=username 且 objectClass=organizationalPerson.
		// username 即我们需要认证的用户名
		//fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", "zhang"),
		filter,
		// 这里是查询返回的属性,以数组形式提供.如果为空则会返回所有的属性
		[]string{},
		nil,
	)

	sr, err := conn.Search(sql)
	if err != nil {
		return nil, err
	}
	dn := sr.Entries[0].DN

	result := &LdapObject{
		DN:       dn,
		Username: username,
		Email:    sr.Entries[0].GetAttributeValue("mail"),
	}
	return result, nil
}
