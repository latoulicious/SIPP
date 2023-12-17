// util/casbin.go

package util

import (
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

var Enforcer *casbin.Enforcer

func InitCasbin() error {
	// Initialize a Xorm adapter with hardcoded credentials
	a, err := xormadapter.NewAdapter("postgres", "user=postgres password=postgres host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		return err
	}

	// Load Casbin model and policy
	Enforcer, err = casbin.NewEnforcer("internal/config/casbin_model.conf", a)
	if err != nil {
		return err
	}

	// Load policies from file
	err = Enforcer.LoadPolicy()
	if err != nil {
		return err
	}

	return nil
}
