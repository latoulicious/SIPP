package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type SoalRepository struct {
	DB *gorm.DB
}

func NewSoalRepository(db *gorm.DB) *SoalRepository {
	return &SoalRepository{DB: db}
}

func (repository *SoalRepository) GetSoal() ([]model.Soal, error) {
	var soal []model.Soal
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("Jurusan").Preload("BankSoal").Find(&soal).Error; err != nil {
		return nil, err
	}
	return soal, nil
}

func (repository *SoalRepository) GetSoalByID(soalID uuid.UUID) (*model.Soal, error) {
	var soal model.Soal
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("Jurusan").Preload("BankSoal").First(&soal, "id = ?", soalID).Error; err != nil {
		return nil, err
	}
	return &soal, nil
}

func (repository *SoalRepository) CreateSoal(soal *model.Soal) error {
	soal.ID = uuid.New()
	if err := repository.DB.Create(soal).Error; err != nil {
		return err
	}
	return repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("Jurusan").Preload("BankSoal").First(soal, "id = ?", soal.ID).Error
}

func (repository *SoalRepository) UpdateSoal(soal *model.Soal) error {
	log.Printf("Updating Soal with ID: %s\n", soal.ID)

	// Log the values of the related entities
	log.Printf("User: %+v\n", soal.User)
	log.Printf("Mapel: %+v\n", soal.Mapel)
	log.Printf("Kelas: %+v\n", soal.Kelas)
	log.Printf("Jurusan: %+v\n", soal.Jurusan)
	log.Printf("BankSoal: %+v\n", soal.BankSoal)

	if err := repository.DB.Save(soal).Error; err != nil {
		log.Printf("Error saving soal: %+v\n", err)
		return err
	}
	return nil
}

func (repository *SoalRepository) DeleteSoal(soalID uuid.UUID) (*model.Soal, error) {
	var soal model.Soal
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("BankSoal").First(&soal, "id = ?", soalID).Error; err != nil {
		return nil, err
	}
	if err := repository.DB.Unscoped().Delete(&soal).Error; err != nil {
		return nil, err
	}
	return &soal, nil
}
