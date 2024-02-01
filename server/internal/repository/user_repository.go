package repository

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repository *UserRepository) GetUsers() ([]model.Users, error) {
	var users []model.Users
	if err := repository.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *UserRepository) GetUserByID(userID uuid.UUID) (*model.Users, error) {
	var user model.Users
	if err := repository.DB.First(&user, "id = ?", userID).Error; err != nil {
		// Log an error if user retrieval fails
		log.Printf("Error fetching user by ID %s: %s\n", userID.String(), err.Error())
		return nil, err
	}

	// Log successful user retrieval
	if os.Getenv("ENVIRONMENT") == "development" {
		log.Printf("User fetched by ID %s: Name: %s\n Role: %s", userID.String(), user.Name, user.Role)
	}
	return &user, nil
}

func (repository *UserRepository) GetUsersPublic() ([]model.Users, error) {
	// Implement logic to fetch all users without requiring JWT authentication
	var users []model.Users
	if err := repository.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *UserRepository) CreateUser(user *model.Users) error {
	user.ID = uuid.New()
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword // Update to store the hashed password
	return repository.DB.Create(user).Error
}

func (repository *UserRepository) UpdateUser(user *model.Users) error {
	return repository.DB.Save(user).Error
}

func (repository *UserRepository) DeleteUser(userID uuid.UUID) error {
	// Instead of using Delete method, use Unscoped().Delete to perform a hard delete
	return repository.DB.Unscoped().Delete(&model.Users{}, "id = ?", userID).Error
}

func (repository *UserRepository) FindByUsername(username string) (*model.Users, error) {
	var user model.Users
	if err := repository.DB.First(&user, "username = ?", username).Error; err != nil {
		log.Printf("Error fetching user by username %s: %s\n", username, err.Error())
		return nil, err
	}
	return &user, nil
}

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (repository *UserRepository) ChangePassword(userID uuid.UUID, newPassword string) error {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return err
	}

	hashedPassword, err := HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return repository.DB.Save(user).Error
}
