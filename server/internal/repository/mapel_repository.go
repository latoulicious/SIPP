package repository

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type MapelRepository struct {
	DB *gorm.DB
}

func NewMapelRepository(db *gorm.DB) *MapelRepository {
	return &MapelRepository{DB: db}
}

func (repository *MapelRepository) GetMapel() ([]model.Mapel, error) {
	var mapel []model.Mapel
	if err := repository.DB.Find(&mapel).Error; err != nil {
		return nil, err
	}
	return mapel, nil
}

func (repository *MapelRepository) GetMapelByID(mapelID uuid.UUID) (*model.Mapel, error) {
	var mapel model.Mapel
	if err := repository.DB.First(&mapel, "id = ?", mapelID).Error; err != nil {
		// Log an error if user retrieval fails
		log.Printf("Error fetching mapel by ID %s: %s\n", mapelID.String(), err.Error())
		return nil, err
	}

	// Log successful user retrieval
	if os.Getenv("ENVIRONMENT") == "development" {
		log.Printf("Mapel fetched by ID %s:", mapelID.String())
	}
	return &mapel, nil
}

func (repository *MapelRepository) CreateMapel(mapel *model.Mapel) error {
	return repository.DB.Create(mapel).Error
}

func (repository *MapelRepository) UpdateMapel(mapelID uuid.UUID, mapel *model.Mapel) error {
	return repository.DB.Model(&model.Mapel{}).Where("id = ?", mapelID).Updates(mapel).Error
}

func (repository *MapelRepository) DeleteMapel(mapelID uuid.UUID) error {
	return repository.DB.Delete(&model.Mapel{}, "id = ?", mapelID).Error
}
