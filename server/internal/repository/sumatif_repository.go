package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type SumatifRepository struct {
	DB *gorm.DB
}

func NewSumatifRepository(db *gorm.DB) *SumatifRepository {
	return &SumatifRepository{DB: db}
}

func (repository *SumatifRepository) GetSumatif() ([]model.Sumatif, error) {
	var sumatif []model.Sumatif
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Preload("BankSoal").Find(&sumatif).Error; err != nil {
		return nil, err
	}
	return sumatif, nil
}

func (repository *SumatifRepository) GetSumatifByID(sumatifID uuid.UUID) (*model.Sumatif, error) {
	var sumatif model.Sumatif
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Preload("BankSoal").First(&sumatif, "id = ?", sumatifID).Error; err != nil {
		return nil, err
	}
	return &sumatif, nil
}

func (repository *SumatifRepository) CreateSumatif(sumatif *model.Sumatif) error {
	// Check if user_id exists in users table
	var user model.Users
	if err := repository.DB.First(&user, "id = ?", sumatif.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with id %s does not exist", sumatif.UserID)
		}
		return err
	}

	// Check if bank_soal_id exists in bank_soals table
	var bankSoal model.BankSoal
	if err := repository.DB.First(&bankSoal, "id = ?", sumatif.BankSoalID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("bank soal with id %s does not exist", sumatif.BankSoalID)
		}
		return err
	}

	// Unmarshal the DynamicFields JSON into a slice of maps
	var dynamicFields []map[string]interface{}
	if err := json.Unmarshal(sumatif.DynamicFields, &dynamicFields); err != nil {
		return err
	}

	// Extract the labels from the DynamicFields array
	var dynamicFieldsLabels []string
	for _, field := range dynamicFields {
		// Access the 'label' key from the map
		if label, ok := field["label"].(string); ok {
			dynamicFieldsLabels = append(dynamicFieldsLabels, label)
		}
	}

	// Assign a new UUID to the sumatif object
	sumatif.ID = uuid.New()

	// Save the sumatif to the database
	if err := repository.DB.Create(&sumatif).Error; err != nil {
		return err
	}

	return nil
}

func (repository *SumatifRepository) UpdateSumatif(sumatif *model.Sumatif) error {
	log.Printf("Updating Sumatif with ID: %s\n", sumatif.ID)

	if err := repository.DB.Save(sumatif).Error; err != nil {
		log.Printf("Error saving sumatif: %+v\n", err)
		return err
	}
	return nil
}

func (repository *SumatifRepository) DeleteSumatif(sumatifID uuid.UUID) (*model.Sumatif, error) {
	var sumatif model.Sumatif
	if err := repository.DB.Unscoped().Where("id = ?", sumatifID).Delete(&sumatif).Error; err != nil {
		return nil, err
	}
	return &sumatif, nil
}
