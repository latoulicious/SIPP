package repository

import (
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
	// Log the formatif object before saving to the database
	log.Printf("Creating Formatif with DynamicFields: %+v\n", formatif.DynamicFields)

	// Assign a new UUID to the formatif object
	formatif.ID = uuid.New()

	// Save the formatif to the database
	if err := repository.DB.Create(&formatif).Error; err != nil {
		return err
	}

	// Log the created formatif object
	log.Printf("Created Formatif object: %+v\n", formatif)

	return nil
}

func (repository *FormatifRepository) UpdateFormatif(formatif *model.Formatif) error {
	log.Printf("Updating Formatif with ID: %s\n", formatif.ID)

	// Assuming formatif.DynamicFields is populated with the updated dynamic fields
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
