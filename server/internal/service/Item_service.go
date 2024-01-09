package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type ItemService struct {
	ItemRepository *repository.ItemRepository
}

func NewItemService(ItemRepository *repository.ItemRepository) *ItemService {
	return &ItemService{
		ItemRepository: ItemRepository,
	}
}

func (service *ItemService) GetItem() ([]model.ItemSoal, error) {
	return service.ItemRepository.GetItem()
}

func (service *ItemService) GetItemByID(itemID uuid.UUID) (*model.ItemSoal, error) {
	return service.ItemRepository.GetItemByID(itemID)
}

func (service *ItemService) CreateItem(item *model.ItemSoal) error {
	return service.ItemRepository.CreateItem(item)
}

func (service *ItemService) UpdateItem(item *model.ItemSoal) error {
	return service.ItemRepository.UpdateItem(item)
}

func (service *ItemService) DeleteItem(itemID uuid.UUID) error {
	return service.ItemRepository.DeleteItem(itemID)
}
