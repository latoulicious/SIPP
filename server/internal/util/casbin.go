// util/casbin.go

package util

import (
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

var Enforcer *casbin.Enforcer

// InitCasbin initializes a Casbin enforcer using a Postgres database.
// It creates a Xorm adapter to connect to Postgres, loads the Casbin
// model and policy files, and loads the policies from file into the
// enforcer. Returns any errors. This allows the rest of the app to
// enforce authorization policies loaded into Casbin.

func InitCasbin() error {
	// Initialize a Xorm adapter with hardcoded credentials
	a, err := xormadapter.NewAdapter("postgres", "user=postgres password=postgres host=localhost port=5432 sslmode=disable")
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
