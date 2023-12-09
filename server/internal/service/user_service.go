// service/user_service.go

package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/controller"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

// GetUsers retrieves a list of users
func (service *UserService) GetUsers() ([]model.Users, error) {
	return service.UserRepository.GetUsers()
}

// GetUserByID retrieves a user by ID
func (service *UserService) GetUserByID(userID uuid.UUID) (*model.Users, error) {
	return service.UserRepository.GetUserByID(userID)
}

// CreateUser creates a new user
func (service *UserService) CreateUser(user *model.Users) error {
	return service.UserRepository.CreateUser(user)
}

// UpdateUser updates an existing user
func (service *UserService) UpdateUser(user *model.Users) error {
	return service.UserRepository.UpdateUser(user)
}

// DeleteUser deletes a user by ID
func (service *UserService) DeleteUser(userID uuid.UUID) error {
	return service.UserRepository.DeleteUser(userID)
}

// Authenticate validates user credentials and returns a JWT token
func (service *UserService) Authenticate(username, password string) (string, error) {
	// Assume userRepo is a repository for user data
	user, err := service.UserRepository.FindByUsername(username)
	if err != nil {
		logrus.Errorf("Error retrieving user from the database: %v", err)
		return "", err
	}

	// Check if the user exists
	if user == nil {
		logrus.Infof("User not found: %s", username)
		return "", errors.New("authentication failed")
	}

	// Check if the provided password matches the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logrus.Infof("Authentication failed for user: %s", username)
		return "", errors.New("authentication failed")
	}

	// Generate and return a JWT token here
	token, err := controller.GenerateJWT(username) // Updated import statement
	if err != nil {
		return "", err
	}
	return token, nil
}
