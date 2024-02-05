package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type RemedialRepository struct {
	DB *gorm.DB
}

func NewRemedialRepository(db *gorm.DB) *RemedialRepository {
	return &RemedialRepository{DB: db}
}

func (repository *RemedialRepository) GetRemedial() ([]model.Remedial, error) {
	var remedial []model.Remedial
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Find(&remedial).Error; err != nil {
		return nil, err
	}
	return remedial, nil
}

func (repository *RemedialRepository) GetRemedialByID(remedialID uuid.UUID) (*model.Remedial, error) {
	var remedial model.Remedial
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").First(&remedial, "id = ?", remedialID).Error; err != nil {
		return nil, err
	}
	return &remedial, nil
}

func (repository *RemedialRepository) CreateRemedial(remedial *model.Remedial) error {
	// Log the remedial object before saving to the database
	log.Printf("Creating Formatif with DynamicFields: %+v\n", remedial.DynamicFields)

	// Assign a new UUID to the remedial object
	remedial.ID = uuid.New()

	// Save the remedial to the database
	if err := repository.DB.Create(&remedial).Error; err != nil {
		return err
	}

	// Log the created remedial object
	log.Printf("Created Remedial object: %+v\n", remedial)

	return nil
}
func (repository *RemedialRepository) UpdateRemedial(remedial *model.Remedial) error {
	log.Printf("Updating Remedial with ID: %s\n", remedial.ID)

	if err := repository.DB.Save(remedial).Error; err != nil {
		log.Printf("Error saving remedial: %+v\n", err)
		return err
	}
	return nil
}

func (repository *RemedialRepository) DeleteRemedial(remedialID uuid.UUID) (*model.Remedial, error) {
	var remedial model.Remedial
	if err := repository.DB.Unscoped().Where("id = ?", remedialID).Delete(&remedial).Error; err != nil {
		return nil, err
	}
	return &remedial, nil
}
