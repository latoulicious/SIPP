package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	// Load the Casbin model and policy from files.
	modelFile := "internal/config/casbin_model.conf"
	policyFile := "internal/config/casbin_policy.csv"

	e, err := casbin.NewEnforcer(modelFile, policyFile)
	if err != nil {
		log.Fatalf("Failed to initialize Casbin: %v", err)
	}

	// Log the loaded policies.
	policies := e.GetPolicy()
	fmt.Println("Loaded Policies:")
	for _, p := range policies {
		fmt.Printf("%v : %v\n", p[0], p[1:])
	}

	// Debug logs for each request.
	sub := "Admin"
	obj := "/api/users/3cdf4f82-85e3-4429-a5de-37c8950942e9"
	act := "PUT"

	fmt.Printf("Request: %s, %s, %s ---> ", sub, obj, act)
	result, err := e.Enforce(sub, obj, act)
	if err != nil {
		fmt.Printf("Casbin Enforce Failed: %v\n", err)
	} else {
		fmt.Printf("Casbin Enforce Result: %v\n", result)
	}
}
