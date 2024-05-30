package main

import (
	"fmt"

	"github.com/latoulicious/SIPP/internal/util"
)

func main() {
	secretKey, err := util.GenerateJWTSecretKey()
	if err != nil {
		fmt.Println("Error generating JWT secret key:", err)
		return
	}
	fmt.Println("Generated JWT Secret Key:", secretKey)
}
