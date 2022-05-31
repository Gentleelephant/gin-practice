package controller

import (
	"crypto/tls"
	"fmt"
	"gin-practice/src/config"
	"gin-practice/src/entity"
	"gin-practice/src/model"
	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap/v3"
	"log"
	"net/http"
)

func Login(c *gin.Context) {

	db := config.GolbalConfig.DB
	dbUser := &entity.User{}

	loginUser := &model.LoginUser{}

	err := c.BindJSON(loginUser)
	if err != nil {
		c.JSON(http.StatusOK, entity.CustomResp{
			Code: 4001,
			Msg:  "参数错误",
			Data: err.Error(),
		})
		c.Abort()
		return
	}

	ldapConfig := config.GolbalConfig.LDAP

	if ldapConfig.Enabled {

		//dial, err1 := ldap.Dial("tcp", fmt.Sprintf("%s:%s", ldapConfig.Host, ldapConfig.Port))
		dial, err1 := ldap.DialURL(fmt.Sprintf("ldap://%s:%s", ldapConfig.Host, ldapConfig.Port))
		if err1 != nil {
			log.Println("连接LDAP服务器失败:", err1)
		}
		defer dial.Close()

		// Reconnect with TLS
		// 建立 StartTLS 连接,这是建立纯文本上的 TLS 协议,允许您将非加密的通讯升级为 TLS 加密而不需要另外使用一个新的端口.
		// 邮件的 POP3 ,IMAP 也有支持类似的 StartTLS,这些都是有 RFC 的

		err1 = dial.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err1 != nil {
			log.Println(err1)
		}

		// First bind with a read only user
		// 先用我们的 bind 账号给 bind 上去
		err1 = dial.Bind(ldapConfig.ManagerDN, ldapConfig.ManagerPassword)
		if err1 != nil {
			log.Println("Bind err: ", err)
		} else {
			log.Println("Bind OK")
		}

		filter := fmt.Sprintf("(&(objectClass=organizationalPerson)(%s=%s))", ldapConfig.LoginAttribute, loginUser.Username)
		log.Println("filter: ", filter)

		sql := ldap.NewSearchRequest(
			// 这里是 basedn,我们将从这个节点开始搜索
			"dc=example,dc=com",
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

		sr, err1 := dial.Search(sql)
		if err1 != nil {
			log.Println("查询失败:", err1)
		}

		if len(sr.Entries) != 1 {
			log.Println("User does not exist or too many entries returned")
			c.JSON(http.StatusOK, entity.CustomResp{
				Code: 4006,
				Msg:  "用户不存在或者查询结果过多",
				Data: nil,
			})
			return
		}

		// 如果没有意外,那么我们就可以获取用户的实际 DN 了
		userdn := sr.Entries[0].DN
		log.Println("userdn", userdn)

		err = dial.Bind(userdn, loginUser.Password)
		if err1 != nil {
			log.Println("Password verification failure", err)
			c.JSON(http.StatusOK, entity.CustomResp{
				Code: 4001,
				Msg:  "用户名或密码错误",
				Data: err1.Error(),
			})
			return
		}

		//dbUser.Username = loginUser.Username
		//dbUser.Email = sr.Entries[0].GetAttributeValue("mail")

		log.Println("dbUser", dbUser)

		result := &entity.User{}

		db.Where(&entity.User{Username: loginUser.Username}, "username").First(result)
		if result.Username == "" {

			result.Username = loginUser.Username
			result.Password = loginUser.Password
			result.Email = sr.Entries[0].GetAttributeValue("mail")

			db.Create(result)
			c.JSON(http.StatusOK, entity.CustomResp{
				Code: 4002,
				Msg:  "用户不存在,重新创建用户",
				Data: result,
			})
			c.Abort()
			return
		}

	} else {

		db.Where(&entity.User{Username: loginUser.Username}, "username").Find(&dbUser)
		if dbUser.Username == "" {
			c.JSON(http.StatusOK, entity.CustomResp{
				Code: 4002,
				Msg:  "用户不存在",
				Data: "",
			})
			c.Abort()
			return
		}

		if dbUser.Password != loginUser.Password {
			c.JSON(http.StatusOK, entity.CustomResp{
				Code: 4003,
				Msg:  "密码错误",
				Data: "",
			})
			c.Abort()
			return
		}

	}
	c.JSON(http.StatusOK, entity.CustomResp{
		Code: 2000,
		Msg:  "登录成功",
		Data: dbUser,
	})
	c.SetCookie("token", "token", 120, "/", c.GetHeader("Host"), false, true)

}
