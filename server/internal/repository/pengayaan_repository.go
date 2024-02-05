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

func (repository *PengayaanRepository) CreatePengayaan(formatif *model.Pengayaan) error {

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
