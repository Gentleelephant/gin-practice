package rbac

import (
	"github.com/casbin/casbin/v2"
	_ "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	_ "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// 从文本中加载模型

var Enforce *casbin.Enforcer

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

	a, err := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/gorm", true)
	if err != nil {
		return
	}
	// 创建执行者。
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return
	}
	Enforce = e
}
