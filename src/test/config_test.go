package test

//
//import (
//	"fmt"
//	"gin-practice/src/config"
//	"gin-practice/src/dao"
//	"gin-practice/src/entity"
//	"testing"
//)
//
//func TestLoadConfig(t *testing.T) {
//
//	err := config.LoadConfig("C:\\work\\code\\goPro\\gin-practice\\src\\config\\config.yaml")
//	if err != nil {
//		t.Error(err)
//	}
//	fmt.Print(config.GlobalConfig)
//
//}
//
//func TestInitDB(t *testing.T) {
//
//	dao.InitDB()
//	db := config.DB
//	user := entity.User{}
//	db.Where(&entity.User{Username: "test3", Password: "test"}, "username", "password").Find(&user)
//	fmt.Print(user)
//
//}

//func TestLDAP(t *testing.T) {
//
//	user := model.LoginUser{
//		Username: "test",
//		Password: "test",
//	}
//
//	dao.InitDB()
//	ldapConfig := config.GlobalConfig.LDAP
//	dial, err1 := ldap.Dial("tcp", fmt.Sprintf("%s:%s", ldapConfig.Host, ldapConfig.Port))
//	if err1 != nil {
//		log.Println("连接LDAP服务器失败:", err1)
//	}
//	defer dial.Close()
//
//	// Reconnect with TLS
//	// 建立 StartTLS 连接,这是建立纯文本上的 TLS 协议,允许您将非加密的通讯升级为 TLS 加密而不需要另外使用一个新的端口.
//	// 邮件的 POP3 ,IMAP 也有支持类似的 StartTLS,这些都是有 RFC 的
//	err1 = dial.StartTLS(&tls.Config{InsecureSkipVerify: true})
//	if err1 != nil {
//		log.Println(err1)
//	}
//
//	// First bind with a read only user
//	// 先用我们的 bind 账号给 bind 上去
//	err1 = dial.Bind(ldapConfig.ManagerDN, ldapConfig.ManagerPassword)
//	if err1 != nil {
//		log.Println("Bind err: ", err1)
//	} else {
//		log.Println("Bind OK")
//	}
//
//	sql := ldap.NewSearchRequest(
//		// 这里是 basedn,我们将从这个节点开始搜索
//		"dc=example,dc=com",
//		// 这里几个参数分别是 scope, derefAliases, sizeLimit, timeLimit,  typesOnly
//		// 详情可以参考 RFC4511 中的定义,文末有链接
//		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
//		// 这里是 LDAP 查询的 Filter.这个例子例子,我们通过查询 uid=username 且 objectClass=organizationalPerson.
//		// username 即我们需要认证的用户名
//		//fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", "zhang"),
//		fmt.Sprintf("(&(objectClass=organizationalPerson)(|(%s=%s)(%s=%s)))", ldapConfig.LoginAttribute, user.Username, ldapConfig.LoginAttribute, user.Username),
//		// 这里是查询返回的属性,以数组形式提供.如果为空则会返回所有的属性
//		[]string{"dn"},
//		nil,
//	)
//
//	sr, err1 := dial.Search(sql)
//	if err1 != nil {
//		log.Println("查询失败:", err1)
//	}
//
//	if len(sr.Entries) != 1 {
//		log.Println("User does not exist or too many entries returned")
//	}
//
//	// 如果没有意外,那么我们就可以获取用户的实际 DN 了
//	userdn := sr.Entries[0].DN
//	log.Println("userdn", userdn)
//
//	err1 = dial.Bind(userdn, user.Password)
//
//}

//func TestADDUser(t *testing.T) {
//
//	// 初始化
//
//	dao.InitDB()
//
//	user := &entity.User{
//		Username: "test",
//		Password: "test",
//		Email:    "1665400978@qq.com",
//		Phone:    "18888888888",
//	}
//	t.Log(user)
//
//	db := config.DB
//	db.Create(user)
//
//}
