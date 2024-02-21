package repository

import (
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
	// Log the pengayaan object before saving to the database
	log.Printf("Creating Pengayaan with DynamicFields: %+v\n", pengayaan.DynamicFields)

	// Assign a new UUID to the pengayaan object
	pengayaan.ID = uuid.New()

	// Save the pengayaan to the database
	if err := repository.DB.Create(&pengayaan).Error; err != nil {
		return err
	}

	// Log the created pengayaan object
	log.Printf("Created Pengayaan object: %+v\n", pengayaan)

	return nil
}

func (repository *PengayaanRepository) UpdatePengayaan(pengayaan *model.Pengayaan) error {
	log.Printf("Updating Pengayaan with ID: %s\n", pengayaan.ID)

	// Assuming pengayaan.DynamicFields is populated with the updated dynamic fields
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
