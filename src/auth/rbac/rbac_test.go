package rbac

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		fmt.Printf("%s can %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s can not %s %s\n", sub, act, obj)
	}
}

func TestRBAC(t *testing.T) {

	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	check(e, "test", "data", "read")

}

//func check2(e *casbin.Enforcer, sub, obj, act string) {
//	ok, _ := e.Enforce(sub, obj, act)
//	if ok {
//		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
//	} else {
//		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
//	}
//}
//
func TestRBAC2(t *testing.T) {
	a, _ := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/gorm", true)
	e, _ := casbin.NewEnforcer("model.conf", a)

	err := e.LoadPolicy()
	if err != nil {
		return
	}

	// Check the permission.
	//check(e, "dajun", "data1", "read")
	//check(e, "lizi", "data2", "write")
	//check(e, "dajun", "data1", "write")
	//check(e, "dajun", "data2", "read")
	check(e, "test", "/v1/test", "GET")
}
