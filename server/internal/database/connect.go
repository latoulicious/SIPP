// database/connect.go

package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/latoulicious/SIPP/internal/config"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DSN string // change from 'dsn' to 'DSN' to avoid shadowing

func ConnectDB() error {
	var err error
	p := config.GetEnv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Failed to parse database port")
		return err
	}

	DSN = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.GetEnv("DB_HOST"), port, config.GetEnv("DB_USER"), config.GetEnv("DB_PASSWORD"), config.GetEnv("DB_NAME"))
	fmt.Println("Connecting to Database with DSN:", DSN)

	DB, err = gorm.Open(postgres.Open(DSN))

	if err != nil {
		log.Println("Failed to connect to the database")
		return err
	}

	fmt.Println("Connection Opened to Database")

	// Auto Migrate Models
	err = DB.AutoMigrate(&model.Users{})
	if err != nil {
		log.Println("Failed to auto migrate models:", err)
		return err
	}
	fmt.Println("Database Migrated")
	return nil
}
