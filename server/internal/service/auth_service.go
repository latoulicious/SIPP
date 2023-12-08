// service/auth_service.go

package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/latoulicious/SIPP/internal/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
	}
}

// Authenticate validates user credentials and returns a JWT token
func (service *AuthService) Authenticate(username, password string) (string, error) {
	// Log the start of the authentication process
	logrus.Infof("Authenticating user: %s", username)

	// Assume userRepo is a repository for user data
	user, err := service.UserRepo.FindByUsername(username)
	if err != nil {
		logrus.Errorf("Error retrieving user from the database: %v", err)
		return "", err
	}

	// Log user information for debugging purposes
	logrus.Infof("User retrieved from the database: %+v", user)

	// Check if the user exists
	if user == nil {
		// Log user not found
		logrus.Infof("User not found: %s", username)
		return "", errors.New("authentication failed")
	}

	// Log the stored password for debugging purposes
	logrus.Infof("Stored password: %s", user.Password)

	// Log the provided password for debugging purposes
	logrus.Infof("Provided password: %s", password)

	// Check if the provided password matches the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// Passwords don't match
		// Log authentication failure
		logrus.Infof("Authentication failed for user: %s", username)
		return "", errors.New("authentication failed")
	}

	// Log successful authentication
	logrus.Infof("Authentication successful for user: %s", username)

	// Generate and return a JWT token here
	token := generateJWT(username)
	return token, nil
}

// comparePasswords compares a stored hashed password with a plain password
func comparePasswords(storedPassword, plainPassword string) bool {
	// Check if the stored password starts with the bcrypt identifier
	if strings.HasPrefix(storedPassword, "$2a$") {
		// The stored password is hashed
		err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(plainPassword))
		return err == nil
	}

	// The stored password is plain, hash the plain password and compare
	hashedPlainPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing plain password:", err)
		return false
	}

	return storedPassword == string(hashedPlainPassword)
}

func generateJWT(username string) string {
	// Example JWT token generation using jwt-go
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Read the JWT secret key from an environment variable
	secretKey := []byte(os.Getenv("JWT_SECRET")) // Replace with os.Getenv("YOUR_JWT_SECRET_KEY")

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("Error generating JWT:", err)
		return ""
	}

	// Log the generated token
	log.Println("Generated JWT Token for user", username, ":", tokenString)

	fmt.Println(tokenString)

	return tokenString
}
