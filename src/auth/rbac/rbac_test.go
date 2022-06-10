package rbac

import (
	"fmt"
	"github.com/casbin/casbin/v2"
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

	check(e, "dajun", "data", "read")
	check(e, "dajun", "data", "write")
	check(e, "lizi", "data", "read")
	check(e, "lizi", "data", "write")
	check(e, "zhangpeng", "data", "delete")

}
