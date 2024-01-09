package repository

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (repository *ItemRepository) GetItem() ([]model.ItemSoal, error) {
	var item []model.ItemSoal
	if err := repository.DB.Find(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (repository *ItemRepository) GetItemByID(itemID uuid.UUID) (*model.ItemSoal, error) {
	var item model.ItemSoal
	if err := repository.DB.First(&item, "id = ?", itemID).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (repository *ItemRepository) CreateItem(item *model.ItemSoal) error {
	item.ID = uuid.New()
	return repository.DB.Create(item).Error
}

func (repository *ItemRepository) UpdateItem(item *model.ItemSoal) error {
	return repository.DB.Save(item).Error
}

func (repository *ItemRepository) DeleteItem(itemID uuid.UUID) error {
	return repository.DB.Delete(&model.ItemSoal{}, "id = ?", itemID).Error
}
