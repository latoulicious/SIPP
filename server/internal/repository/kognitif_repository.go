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

type KognitifRepository struct {
	DB *gorm.DB
}

func NewKognitifRepository(db *gorm.DB) *KognitifRepository {
	return &KognitifRepository{DB: db}
}

func (repository *KognitifRepository) GetKognitif() ([]model.Kognitif, error) {
	var kognitif []model.Kognitif
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Preload("BankSoal").Find(&kognitif).Error; err != nil {
		return nil, err
	}
	return kognitif, nil
}

func (repository *KognitifRepository) GetKognitifByID(kognitifID uuid.UUID) (*model.Kognitif, error) {
	var kognitif model.Kognitif
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Preload("BankSoal").First(&kognitif, "id = ?", kognitifID).Error; err != nil {
		return nil, err
	}
	return &kognitif, nil
}

func (repository *KognitifRepository) CreateKognitif(kognitif *model.Kognitif) error {
	// Check if user_id exists in users table
	var user model.Users
	if err := repository.DB.First(&user, "id = ?", kognitif.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with id %s does not exist", kognitif.UserID)
		}
		return err
	}

	// Check if bank_soal_id exists in bank_soals table
	var bankSoal model.BankSoal
	if err := repository.DB.First(&bankSoal, "id = ?", kognitif.BankSoalID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("bank soal with id %s does not exist", kognitif.BankSoalID)
		}
		return err
	}

	// Unmarshal the DynamicFields JSON into a slice of maps
	var dynamicFields []map[string]interface{}
	if err := json.Unmarshal(kognitif.DynamicFields, &dynamicFields); err != nil {
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

	// Assign a new UUID to the kognitif object
	kognitif.ID = uuid.New()

	// Save the kognitif to the database
	if err := repository.DB.Create(&kognitif).Error; err != nil {
		return err
	}

	return nil
}

func (repository *KognitifRepository) UpdateKognitif(kognitif *model.Kognitif) error {
	log.Printf("Updating Kognitif with ID: %s\n", kognitif.ID)

	if err := repository.DB.Save(kognitif).Error; err != nil {
		log.Printf("Error saving kognitif: %+v\n", err)
		return err
	}
	return nil
}

func (repository *KognitifRepository) DeleteKognitif(kognitifID uuid.UUID) (*model.Kognitif, error) {
	var kognitif model.Kognitif
	if err := repository.DB.Unscoped().Where("id = ?", kognitifID).Delete(&kognitif).Error; err != nil {
		return nil, err
	}
	return &kognitif, nil
}
