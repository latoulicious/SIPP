package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/latoulicious/SIPP/internal/config"
	models "github.com/latoulicious/SIPP/internal/model/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connects to the database.
//
// It does not take any parameters.
// It does not return any values.
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Panic("Failed to parse database port env var")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Connected to database!")

	DB.AutoMigrate(&models.Users{})
	fmt.Println("Database Migrated")
}
