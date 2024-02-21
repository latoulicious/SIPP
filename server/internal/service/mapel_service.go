package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type MapelService struct {
	MapelRepository *repository.MapelRepository
}

// NewMapelService creates a new instance of MapelService
func NewMapelService(mapelRepository *repository.MapelRepository) *MapelService {
	return &MapelService{
		MapelRepository: mapelRepository,
	}
}

// GetMapel retrieves a list of mapel
func (service *MapelService) GetMapel() ([]model.Mapel, error) {
	return service.MapelRepository.GetMapel()
}

// GetMapelByID retrieves a mapel by ID
func (service *MapelService) GetMapelByID(id uuid.UUID) (*model.Mapel, error) {
	return service.MapelRepository.GetMapelByID(id)
}

// Public Endpoint
func (service *MapelService) GetMapelPublic() ([]model.Mapel, error) {
	// Implement logic to fetch all mapel without requiring JWT authentication
	return service.MapelRepository.GetMapelPublic()
}

// CreateMapel creates a new mapel
func (service *MapelService) CreateMapel(mapel *model.Mapel) error {
	return service.MapelRepository.CreateMapel(mapel)
}

// UpdateMapel updates an existing mapel
func (service *MapelService) UpdateMapel(id uuid.UUID, mapel *model.Mapel) error {
	return service.MapelRepository.UpdateMapel(id, mapel)
}

// DeleteMapel deletes a mapel by ID
func (service *MapelService) DeleteMapel(id uuid.UUID) error {
	return service.MapelRepository.DeleteMapel(id)
}
