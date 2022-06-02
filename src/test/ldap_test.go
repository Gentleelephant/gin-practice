package test

import (
	"fmt"
	"gin-practice/src/auth"
	"gin-practice/src/config"
	"gin-practice/src/dao"
	"gin-practice/src/entity"
	"github.com/go-ldap/ldap/v3"
	"log"
	"testing"
)

func TestLdap(t *testing.T) {

	err := config.LoadConfig(config.ConfigPath)
	config.DB = dao.InitDB()
	if err != nil {
		return
	}
	provider := auth.GetProvider()
	authentication, dn, _ := provider.Authentication("zhang", "zhang")
	if authentication == false {
		t.Error("authentication failed")
		return
	}
	// 判断表中是否以及包含该用户
	user := &entity.User{
		Username: "lihua",
		Email:    "1132960613@qq.com",
	}

	take := dao.CheckUser(user)
	t.Log(take)
	if take != 1 {
		t.Log("user not exist.......")
		t.Log("insert user......")
		// 插入用户
		err := dao.AddUserAndUserDn(user, &entity.UserDn{
			Username: user.Username,
			Dn:       dn,
		})
		if err != nil {
			return
		}
		return
	}
	t.Log("user exist.......")

}

func TestPool(t *testing.T) {
	err := config.LoadConfig(config.ConfigPath)
	if err != nil {
		return
	}
	pool, err := auth.NewPool(config.GlobalConfig.LDAP.NewLdapClient, 10)
	if err != nil {
		return
	}
	conn, err := pool.Acquire()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("pool success")
	err = conn.Bind(config.GlobalConfig.LDAP.ManagerDN, config.GlobalConfig.LDAP.ManagerPassword)
	if err != nil {
		t.Log("Bind Manger DN error", err)
		return
	}

	username := "zhang"
	// Search for user
	filter := fmt.Sprintf("(&(objectClass=organizationalPerson)(%s=%s))", config.GlobalConfig.LDAP.LoginAttribute, username)
	log.Println("filter: ", filter)

	sql := ldap.NewSearchRequest(
		// 这里是 basedn,我们将从这个节点开始搜索
		config.GlobalConfig.LDAP.UserSearchBase,
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
		t.Log("Query failed:", err)
		return
	}
	dn := sr.Entries[0].DN
	fmt.Println("user dn: ", dn)

}
