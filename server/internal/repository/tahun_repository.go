package repository

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type TahunRepository struct {
	DB *gorm.DB
}

func NewTahunRepository(db *gorm.DB) *TahunRepository {
	return &TahunRepository{DB: db}
}

func (repository *TahunRepository) GetTahun() ([]model.TahunAjar, error) {
	var tahun []model.TahunAjar
	if err := repository.DB.Find(&tahun).Error; err != nil {
		return nil, err
	}
	return tahun, nil
}

func (repository *TahunRepository) GetTahunByID(tahunID uuid.UUID) (*model.TahunAjar, error) {
	var tahun model.TahunAjar
	if err := repository.DB.First(&tahun, "id = ?", tahunID).Error; err != nil {
		// Log an error if tahun retrieval fails
		log.Printf("Error fetching tahun by ID %s: %s\n", tahunID.String(), err.Error())
		return nil, err
	}

	// Log successful tahun retrieval
	if os.Getenv("ENVIRONMENT") == "development" {
		log.Printf("Tahun fetched by ID %s:", tahunID.String())
	}
	return &tahun, nil
}

func (repository *TahunRepository) GetTahunPublic() ([]model.TahunAjar, error) {
	// Implement logic to fetch all tahun without requiring JWT authentication
	var tahun []model.TahunAjar
	if err := repository.DB.Find(&tahun).Error; err != nil {
		return nil, err
	}
	return tahun, nil
}

func (repository *TahunRepository) CreateTahun(tahun *model.TahunAjar) error {
	return repository.DB.Create(tahun).Error
}

func (repository *TahunRepository) UpdateTahun(tahun *model.TahunAjar) error {
	return repository.DB.Save(tahun).Error
}

func (repository *TahunRepository) DeleteTahun(tahunID uuid.UUID) error {
	// Instead of using Delete method, use Unscoped().Delete to perform a hard delete
	return repository.DB.Unscoped().Delete(&model.TahunAjar{}, "id = ?", tahunID).Error
}
