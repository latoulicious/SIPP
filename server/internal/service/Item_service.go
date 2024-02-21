package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type ItemService struct {
	ItemRepository *repository.ItemRepository
}

// NewItemService creates a new instance of ItemService
func NewItemService(ItemRepository *repository.ItemRepository) *ItemService {
	return &ItemService{
		ItemRepository: ItemRepository,
	}
}

// GetItem retrieves a list of items
func (service *ItemService) GetItem() ([]model.ItemSoal, error) {
	return service.ItemRepository.GetItem()
}

// GetItemByID retrieves an item by ID
func (service *ItemService) GetItemByID(itemID uuid.UUID) (*model.ItemSoal, error) {
	return service.ItemRepository.GetItemByID(itemID)
}

// CreateItem creates a new item
func (service *ItemService) CreateItem(item *model.ItemSoal) error {
	return service.ItemRepository.CreateItem(item)
}

// UpdateItem updates an existing item
func (service *ItemService) UpdateItem(item *model.ItemSoal) error {
	return service.ItemRepository.UpdateItem(item)
}

// DeleteItem deletes an item by ID
func (service *ItemService) DeleteItem(itemID uuid.UUID) error {
	return service.ItemRepository.DeleteItem(itemID)
}
