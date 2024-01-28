package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type CapaianRepository struct {
	DB *gorm.DB
}

func NewCapaianRepository(db *gorm.DB) *CapaianRepository {
	return &CapaianRepository{DB: db}
}

func (repository *CapaianRepository) GetCapaian() ([]model.Capaian, error) {
	var capaian []model.Capaian
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Find(&capaian).Error; err != nil {
		return nil, err
	}
	return capaian, nil
}

func (repository *CapaianRepository) GetCapaianByID(capaianID uuid.UUID) (*model.Capaian, error) {
	var capaian model.Capaian
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").First(&capaian, "id = ?", capaianID).Error; err != nil {
		return nil, err
	}
	return &capaian, nil
}

func (repository *CapaianRepository) CreateCapaian(capaian *model.Capaian) error {
	capaian.ID = uuid.New()
	if err := repository.DB.Create(capaian).Error; err != nil {
		return err
	}
	return repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").First(capaian, "id = ?", capaian.ID).Error
}

func (repository *CapaianRepository) UpdateCapaian(capaian *model.Capaian) error {
	log.Printf("Updating Capaian with ID: %s\n", capaian.ID)

	// Log the values of the related entities
	log.Printf("User: %+v\n", capaian.User)
	log.Printf("Mapel: %+v\n", capaian.Mapel)
	log.Printf("Kelas: %+v\n", capaian.Kelas)
	log.Printf("TahunAjar: %+v\n", capaian.TahunAjar)

	if err := repository.DB.Save(capaian).Error; err != nil {
		log.Printf("Error saving capaian: %+v\n", err)
		return err
	}
	return repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").First(capaian, "id = ?", capaian.ID).Error
}

func (repository *CapaianRepository) DeleteCapaian(capaianID uuid.UUID) (*model.Capaian, error) {
	var capaian model.Capaian
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").First(&capaian, "id = ?", capaianID).Error; err != nil {
		return nil, err
	}
	if err := repository.DB.Unscoped().Delete(&capaian).Error; err != nil {
		return nil, err
	}
	return &capaian, nil
}
