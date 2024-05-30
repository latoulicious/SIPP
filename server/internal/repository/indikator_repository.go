package repository

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type IndikatorRepository struct {
	DB *gorm.DB
}

func NewIndikatorRepository(db *gorm.DB) *IndikatorRepository {
	return &IndikatorRepository{DB: db}
}

func (repository *IndikatorRepository) GetIndikator() ([]model.Indikator, error) {
	var indikator []model.Indikator
	if err := repository.DB.Find(&indikator).Error; err != nil {
		return nil, err
	}

	return indikator, nil
}

func (repository *IndikatorRepository) GetIndikatorByID(indikatorID uuid.UUID) (*model.Indikator, error) {
	var indikator model.Indikator
	if err := repository.DB.First(&indikator, "id =?", indikatorID).Error; err != nil {
		return nil, err
	}
	return &indikator, nil
}

func (repository *IndikatorRepository) CreateIndikator(indikator *model.Indikator) error {
	indikator.ID = uuid.New()
	return repository.DB.Create(indikator).Error
}

func (repository *IndikatorRepository) UpdateIndikator(indikator *model.Indikator) error {
	return repository.DB.Save(indikator).Error
}

func (repository *IndikatorRepository) DeleteIndikator(indikatorID uuid.UUID) error {
	if err := repository.DB.Unscoped().Where("id =?", indikatorID).Delete(&model.Indikator{}).Error; err != nil {
		return err
	}

	return nil
}
