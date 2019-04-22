package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/mongodb-adapter"
)

func main() {
	a := mongodbadapter.NewAdapter("127.0.0.1:27017/casbin")

	e := casbin.NewEnforcer("examples/rbac_model.conf", a)

	e.LoadPolicy()

	aliceHasAccess(e)

	e.AddPolicy("alice", "data1", "read")

	aliceHasAccess(e)

	e.SavePolicy()
}

func aliceHasAccess(e *casbin.Enforcer) {
	if e.Enforce("alice", "data1", "read") == true {
		fmt.Println("alice has access read access to data1")
	} else {
		fmt.Println("alice has not access read access to data1")
	}
}

