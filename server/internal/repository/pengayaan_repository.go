package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type PengayaanRepository struct {
	DB *gorm.DB
}

func NewPengayaanRepository(db *gorm.DB) *PengayaanRepository {
	return &PengayaanRepository{DB: db}
}

func (repository *PengayaanRepository) GetPengayaan() ([]model.Pengayaan, error) {
	var pengayaan []model.Pengayaan
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Find(&pengayaan).Error; err != nil {
		return nil, err
	}
	return pengayaan, nil
}

func (repository *PengayaanRepository) GetPengayaanByID(pengayaanID uuid.UUID) (*model.Pengayaan, error) {
	var pengayaan model.Pengayaan
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").First(&pengayaan, "id = ?", pengayaanID).Error; err != nil {
		return nil, err
	}
	return &pengayaan, nil
}

func (repository *PengayaanRepository) CreatePengayaan(pengayaan *model.Pengayaan) error {
	// Check if user_id exists in users table
	var user model.Users
	if err := repository.DB.First(&user, "id = ?", pengayaan.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with id %s does not exist", pengayaan.UserID)
		}
		return err
	}

	pengayaan.ID = uuid.New()
	if err := repository.DB.Create(&pengayaan).Error; err != nil {
		return err
	}
	return nil
}

func (repository *PengayaanRepository) UpdatePengayaan(pengayaan *model.Pengayaan) error {
	log.Printf("Updating Pengayaan with ID: %s\n", pengayaan.ID)

	if err := repository.DB.Save(pengayaan).Error; err != nil {
		log.Printf("Error saving pengayaan: %+v\n", err)
		return err
	}
	return nil
}

func (repository *PengayaanRepository) DeletePengayaan(pengayaanID uuid.UUID) (*model.Pengayaan, error) {
	var pengayaan model.Pengayaan
	if err := repository.DB.Unscoped().Where("id = ?", pengayaanID).Delete(&pengayaan).Error; err != nil {
		return nil, err
	}
	return &pengayaan, nil
}
