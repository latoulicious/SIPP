// migration.go
package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/config"
	"github.com/latoulicious/SIPP/internal/database"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

// main is the entry point for the password migration utility. It loads the
// configuration, connects to the database, retrieves all users with plain text
// passwords, hashes the passwords, updates the database records, and logs
// the result.

func main() {
	// Load configuration from .env file
	config.LoadConfig()

	// Connect to the database
	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Get all users with plain-text passwords
	users, err := getAllUsersWithPlainTextPasswords()
	if err != nil {
		log.Fatal("Failed to retrieve users:", err)
	}

	// Update each user with a hashed password
	for _, user := range users {
		hashedPassword, err := repository.HashPassword(user.Password)
		if err != nil {
			log.Fatal("Failed to hash password:", err)
		}

		// Update the user record in the database
		err = updateUserPassword(user.ID, hashedPassword)
		if err != nil {
			log.Fatal("Failed to update user password:", err)
		}
	}

	log.Println("Password migration completed successfully.")
}

// getAllUsersWithPlainTextPasswords retrieves all users with plain-text passwords from the database.
// It queries the database for all User models and returns them in a slice.
// Returns an error if the database query fails.

func getAllUsersWithPlainTextPasswords() ([]model.Users, error) {
	var users []model.Users
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// updateUserPassword updates the password field in the database for the
// user with the given ID, setting it to the provided hashed password.

func updateUserPassword(userID uuid.UUID, hashedPassword string) error {
	return database.DB.Model(&model.Users{}).Where("id = ?", userID).Update("password", hashedPassword).Error
}
