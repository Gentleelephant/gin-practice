package rbac

import (
	"fmt"
	"gin-practice/src/config"
	"github.com/casbin/casbin/v2"
	_ "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	_ "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// 从文本中加载模型

var Enforcer *casbin.Enforcer

func InitCasbin() {
	// 从字符串初始化模型
	text := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`
	m, err := model.NewModelFromString(text)
	if err != nil {
		return
	}

	mysql := config.GlobalConfig.Mysql

	user := mysql.User
	password := mysql.Password
	host := mysql.Host
	port := mysql.Port
	dbname := mysql.Database

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	fmt.Println(dataSourceName)
	a, err := gormadapter.NewAdapter("mysql", dataSourceName, true)
	if err != nil {
		return
	}
	// 创建执行者。
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return
	}
	Enforcer = e
}
