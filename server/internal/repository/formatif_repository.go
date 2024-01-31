package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type FormatifRepository struct {
	DB *gorm.DB
}

func NewFormatifRepository(db *gorm.DB) *FormatifRepository {
	return &FormatifRepository{DB: db}
}

func (repository *FormatifRepository) GetFormatif() ([]model.Formatif, error) {
	var formatif []model.Formatif
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Find(&formatif).Error; err != nil {
		return nil, err
	}
	return formatif, nil
}

func (repository *FormatifRepository) GetFormatifByID(formatifID uuid.UUID) (*model.Formatif, error) {
	var formatif model.Formatif
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").First(&formatif, "id = ?", formatifID).Error; err != nil {
		return nil, err
	}
	return &formatif, nil
}

func (repository *FormatifRepository) CreateFormatif(formatif *model.Formatif) error {
	// Check if user_id exists in users table
	var user model.Users
	if err := repository.DB.First(&user, "id = ?", formatif.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with id %s does not exist", formatif.UserID)
		}
		return err
	}

	formatif.ID = uuid.New()
	if err := repository.DB.Create(&formatif).Error; err != nil {
		return err
	}
	return nil
}

func (repository *FormatifRepository) UpdateFormatif(formatif *model.Formatif) error {
	log.Printf("Updating Formatif with ID: %s\n", formatif.ID)

	if err := repository.DB.Save(formatif).Error; err != nil {
		log.Printf("Error saving formatif: %+v\n", err)
		return err
	}
	return nil
}

func (repository *FormatifRepository) DeleteFormatif(formatifID uuid.UUID) (*model.Formatif, error) {
	var formatif model.Formatif
	if err := repository.DB.Unscoped().Where("id = ?", formatifID).Delete(&formatif).Error; err != nil {
		return nil, err
	}
	return &formatif, nil
}
