package repository

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type PenilaianRepository struct {
	DB *gorm.DB
}

func NewPenilaianRepository(db *gorm.DB) *PenilaianRepository {
	return &PenilaianRepository{DB: db}
}

func (repository *PenilaianRepository) GetPenilaian() ([]model.Penilaian, error) {
	var penilaian []model.Penilaian
	if err := repository.DB.Find(&penilaian).Error; err != nil {
		return nil, err
	}
	return penilaian, nil
}

func (repository *PenilaianRepository) GetPenilaianByID(penilaianID uuid.UUID) (*model.Penilaian, error) {
	var penilaian model.Penilaian
	if err := repository.DB.First(&penilaian, "id = ?", penilaianID).Error; err != nil {
		return nil, err
	}
	return &penilaian, nil
}

func (repository *PenilaianRepository) CreatePenilaian(penilaian *model.Penilaian) error {
	return repository.DB.Create(penilaian).Error
}

func (repository *PenilaianRepository) UpdatePenilaian(penilaian *model.Penilaian) error {
	return repository.DB.Save(penilaian).Error
}

func (repository *PenilaianRepository) DeletePenilaian(penilaianID uuid.UUID) error {
	// Instead of using Delete method, use Unscoped().Delete to perform a hard delete
	return repository.DB.Unscoped().Delete(&model.Penilaian{}, "id = ?", penilaianID).Error
}
