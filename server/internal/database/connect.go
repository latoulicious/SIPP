// database/connect.go

package database

import (
	"fmt"
	"log"
	"os"
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

	if os.Getenv("ENVIRONMENT") == "development" {
		DSN = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.GetEnv("DB_HOST"), port, config.GetEnv("DB_USER"), config.GetEnv("DB_PASSWORD"), config.GetEnv("DB_NAME"))
		log.Println("Connecting to Database with DSN:", DSN)
	}

	DB, err = gorm.Open(postgres.Open(DSN))

	if err != nil {
		log.Println("Failed to connect to the database")
		return err
	}

	log.Println("Connection Opened to Database")

	// Auto Migrate Models
	err = DB.AutoMigrate(
		&model.Users{},
		&model.TahunAjar{},
		&model.Mapel{},
		&model.Kelas{},
		&model.Jurusan{},
		&model.Capaian{},
		&model.AlurTP{},
		&model.ModulAjar{},
		// &model.Penilaian{},
		&model.BaseModel{},
		&model.Kognitif{},
		&model.Formatif{},
		&model.Sumatif{},
		&model.Pengayaan{},
		&model.Remedial{},
		&model.Soal{},
		&model.ItemSoal{},
		&model.BankSoal{},
	)
	if err != nil {
		log.Println("Failed to auto migrate models:", err)
		return err
	}
	log.Println("Database Migrated")
	return nil
}
