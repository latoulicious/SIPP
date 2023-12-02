package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config loads the value of the specified key from the .env file.
//
// Parameters:
// - key: the key of the value to be retrieved from the .env file
//
// Return type:
// - string: the value associated with the specified key
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	return os.Getenv(key)
}
