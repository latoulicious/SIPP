package repository

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type TingkatRepository struct {
	DB *gorm.DB
}

func NewTingkatRepository(db *gorm.DB) *TingkatRepository {
	return &TingkatRepository{DB: db}
}

func (repository *TingkatRepository) GetIndikatorTingkat() ([]model.IndikatorTingkat, error) {
	var indikatorTingkats []model.IndikatorTingkat
	if err := repository.DB.Find(&indikatorTingkats).Error; err != nil {
		return nil, err
	}
	return indikatorTingkats, nil
}

func (repository *TingkatRepository) GetIndikatorTingkatByID(id uuid.UUID) (*model.IndikatorTingkat, error) {
	var indikatorTingkat model.IndikatorTingkat
	if err := repository.DB.First(&indikatorTingkat, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &indikatorTingkat, nil
}

func (repository *TingkatRepository) CreateIndikatorTingkat(indikatorID uuid.UUID, kesukaranID uuid.UUID) error {
	indikatorTingkat := model.IndikatorTingkat{
		ID:          uuid.New(), // Generate a new UUID here
		IndikatorID: indikatorID,
		KesukaranID: kesukaranID,
	}

	return repository.DB.Create(&indikatorTingkat).Error
}
