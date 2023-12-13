// service/auth_service.go

package service

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/latoulicious/SIPP/internal/controller"
	"github.com/latoulicious/SIPP/internal/repository"
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
	fmt.Printf("Authenticating user: %s", username)

	// Assume userRepo is a repository for user data
	user, err := service.UserRepo.FindByUsername(username)
	if err != nil {
		fmt.Printf("Error retrieving user from the database: %v", err)
		return "", err
	}

	// Log user information for debugging purposes
	fmt.Printf("User retrieved from the database: %+v", user)

	// Check if the user exists
	if user == nil {
		// Log user not found
		fmt.Printf("User not found: %s", username)
		return "", errors.New("authentication failed")
	}

	// Log the stored password for debugging purposes
	fmt.Printf("Stored password: %s", user.Password)

	// Log the provided password for debugging purposes
	fmt.Printf("Provided password: %s", password)

	// Check if the provided password matches the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// Passwords don't match
		// Log authentication failure
		fmt.Printf("Authentication failed for user: %s", username)
		return "", errors.New("authentication failed")
	}

	// Log successful authentication
	fmt.Printf("Authentication successful for user: %s", username)

	// Generate and return a JWT token here
	token, err := controller.GenerateJWT(username) // Updated import statement
	if err != nil {
		return "", err
	}
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
