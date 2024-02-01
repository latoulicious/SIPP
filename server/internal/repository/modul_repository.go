package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type ModulRepository struct {
	DB *gorm.DB
}

func NewModulRepository(db *gorm.DB) *ModulRepository {
	return &ModulRepository{DB: db}
}

func (repository *ModulRepository) GetModul() ([]model.ModulAjar, error) {
	var modul []model.ModulAjar
	if err := repository.DB.Preload("User").Preload("Kelas").Preload("TahunAjar").Find(&modul).Error; err != nil {
		return nil, err
	}
	return modul, nil
}

func (repository *ModulRepository) GetModulByID(modulID uuid.UUID) (model.ModulAjar, error) {
	var modul model.ModulAjar
	if err := repository.DB.Preload("User").Preload("Kelas").Preload("TahunAjar").Find(&modul).Error; err != nil {
		return model.ModulAjar{}, err
	}
	return modul, nil
}

func (repository *ModulRepository) CreateModul(modul *model.ModulAjar) error {
	modul.ID = uuid.New()
	if err := repository.DB.Create(&modul).Error; err != nil {
		return err
	}
	return repository.DB.Preload("User").Preload("Kelas").Preload("TahunAjar").Find(&modul).Error
}

func (repository *ModulRepository) UpdateModul(modul *model.ModulAjar) error {
	log.Printf("Updating Modul with ID: %s\n", modul.ID)

	// Log the values of the related entities
	log.Printf("User: %+v\n", modul.User)
	log.Printf("Kelas: %+v\n", modul.Kelas)
	log.Printf("TahunAjar: %+v\n", modul.TahunAjar)

	if err := repository.DB.Save(modul).Error; err != nil {
		log.Printf("Error saving modul: %+v\n", err)
		return err
	}
	return nil
}

func (repository *ModulRepository) DeleteModul(modulID uuid.UUID) (*model.ModulAjar, error) {
	var modul model.ModulAjar
	if err := repository.DB.Preload("User").Preload("Kelas").Preload("TahunAjar").First(&modul, "id = ?", modulID).Error; err != nil {
		return nil, err
	}
	if err := repository.DB.Unscoped().Delete(&modul).Error; err != nil {
		return nil, err
	}
	return &modul, nil
}
