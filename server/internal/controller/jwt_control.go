// controller/jwt_control.go

package controller

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	tokenMap  = make(map[string]string)
	tokenLock sync.Mutex
)

// GenerateJWT generates a new JWT token for the specified user. It takes in the user ID,
// username, name and role. It first checks if a token already exists for the username,
// returning the existing token if found. Otherwise it generates a new JWT token using the
// jwt-go library, sets claims including the user info, signs it using the JWT secret read
// from the environment, and stores the token in a map before returning it. In development
// environment, it also logs the generated token.

func GenerateJWT(userID, username, name, role string) (string, error) {
	tokenLock.Lock()
	defer tokenLock.Unlock()

	// Check if a token already exists for the given username
	if existingToken, ok := tokenMap[username]; ok {
		return existingToken, nil
	}

	// Example JWT token generation using jwt-go
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID // Include the user's ID in the claims
	claims["username"] = username
	claims["name"] = name
	claims["role"] = role                                 // Include the user's role in the claims
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Read the JWT secret key from an environment variable
	secretKey := []byte(os.Getenv("JWT_SECRET")) // Replace with os.Getenv("YOUR_JWT_SECRET_KEY")

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("Error generating JWT:", err)
		return "", err
	}

	// Log the generated token only in development environment
	if os.Getenv("ENVIRONMENT") == "development" {
		log.Println("Generated JWT Token for user", username, ":", tokenString)
	}

	// Store the generated token in the map
	tokenMap[username] = tokenString

	return tokenString, nil
}
