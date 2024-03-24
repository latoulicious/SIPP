package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type KesukaranService struct {
	KesukaranRepository *repository.KesukaranRepository
}

// NewKesukaranService creates a new instance of KesukaranService
func NewKesukaranService(kesukaranRepository *repository.KesukaranRepository) *KesukaranService {
	return &KesukaranService{
		KesukaranRepository: kesukaranRepository,
	}
}

// GetKesukaran retrieve a list of kesukaran
func (service *KesukaranService) GetKesukaran() ([]model.Kesukaran, error) {
	return service.KesukaranRepository.GetKesukaran()
}

// GetKesukaranByID retrieve a kesukaran by id
func (service *KesukaranService) GetKesukaranByID(kesukaranID uuid.UUID) (*model.Kesukaran, error) {
	return service.KesukaranRepository.GetKesukaranByID(kesukaranID)
}

// CreateKesukaran creates a new kesukaran
func (service *KesukaranService) CreateKesukaran(kesukaran *model.Kesukaran) error {
	return service.KesukaranRepository.CreateKesukaran(kesukaran)
}

// UpdateKesukaran updates an existing kesukaran
func (service *KesukaranService) UpdateKesukaran(kesukaran *model.Kesukaran) error {
	return service.KesukaranRepository.UpdateKesukaran(kesukaran)
}

// DeleteKesukaran deletes a kesukaran by ID
func (service *KesukaranService) DeleteKesukaran(kesukaranID uuid.UUID) error {
	return service.KesukaranRepository.DeleteKesukaran(kesukaranID)
}
