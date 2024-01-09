package repository

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type SoalRepository struct {
	DB *gorm.DB
}

func NewSoalRepository(db *gorm.DB) *SoalRepository {
	return &SoalRepository{DB: db}
}

func (repository *SoalRepository) GetSoal() ([]model.Soal, error) {
	var soal []model.Soal
	if err := repository.DB.Find(&soal).Error; err != nil {
		return nil, err
	}
	return soal, nil
}

func (repository *SoalRepository) GetSoalByID(soalID uuid.UUID) (*model.Soal, error) {
	var soal model.Soal
	if err := repository.DB.First(&soal, "id = ?", soalID).Error; err != nil {
		return nil, err
	}
	return &soal, nil
}

func (repository *SoalRepository) CreateSoal(soal *model.Soal) error {
	soal.ID = uuid.New()
	return repository.DB.Create(soal).Error
}

func (repository *SoalRepository) UpdateSoal(soal *model.Soal) error {
	return repository.DB.Save(soal).Error
}

func (repository *SoalRepository) DeleteSoal(soalID uuid.UUID) error {
	return repository.DB.Delete(&model.Soal{}, "id = ?", soalID).Error
}
