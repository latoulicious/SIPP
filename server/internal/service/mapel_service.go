package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type MapelService struct {
	mapelRepository *repository.MapelRepository
}

func NewMapelService(mapelRepository *repository.MapelRepository) *MapelService {
	return &MapelService{mapelRepository: mapelRepository}
}

func (service *MapelService) GetMapel() ([]model.Mapel, error) {
	return service.mapelRepository.GetMapel()
}

func (service *MapelService) GetMapelByID(id uuid.UUID) (*model.Mapel, error) {
	return service.mapelRepository.GetMapelByID(id)
}

func (service *MapelService) CreateMapel(mapel *model.Mapel) error {
	return service.mapelRepository.CreateMapel(mapel)
}

func (service *MapelService) UpdateMapel(id uuid.UUID, mapel *model.Mapel) error {
	return service.mapelRepository.UpdateMapel(id, mapel)
}

func (service *MapelService) DeleteMapel(id uuid.UUID) error {
	return service.mapelRepository.DeleteMapel(id)
}
