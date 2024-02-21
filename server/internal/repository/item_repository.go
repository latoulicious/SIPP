package repository

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type ItemSoalRepository struct {
	DB *gorm.DB
}

func NewItemSoalRepository(db *gorm.DB) *ItemSoalRepository {
	return &ItemSoalRepository{DB: db}
}

func (repository *ItemSoalRepository) GetItemSoal() ([]model.ItemSoal, error) {
	var itemsoal []model.ItemSoal
	if err := repository.DB.Find(&itemsoal).Error; err != nil {
		return nil, err
	}

	return itemsoal, nil
}

func (repository *ItemSoalRepository) GetItemSoalByID(itemsoalID uuid.UUID) (*model.ItemSoal, error) {
	var itemsoal model.ItemSoal
	if err := repository.DB.First(&itemsoal, "id = ?", itemsoalID).Error; err != nil {
		return nil, err
	}
	return &itemsoal, nil
}
