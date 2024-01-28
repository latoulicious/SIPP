package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type CapaianService struct {
	CapaianRepository *repository.CapaianRepository
}

// NewCapaianService creates a new instance of CapaianService
func NewCapaianService(capaianRepository *repository.CapaianRepository) *CapaianService {
	return &CapaianService{
		CapaianRepository: capaianRepository,
	}
}

// GetCapaian retrieve a list of capaian
func (service *CapaianService) GetCapaian() ([]model.Capaian, error) {
	return service.CapaianRepository.GetCapaian()
}

// GetCapaianByID retrieve a capaian by id
func (service *CapaianService) GetCapaianByID(capaianID uuid.UUID) (*model.Capaian, error) {
	return service.CapaianRepository.GetCapaianByID(capaianID)
}

// CreateCapaian creates a new capaian
func (service *CapaianService) CreateCapaian(capaian *model.Capaian) error {
	return service.CapaianRepository.CreateCapaian(capaian)
}

// UpdateCapaian updates an existing capaian
func (service *CapaianService) UpdateCapaian(capaian *model.Capaian) error {
	return service.CapaianRepository.UpdateCapaian(capaian)
}

// DeleteCapaian deletes a capaian by ID
func (service *CapaianService) DeleteCapaian(capaianID uuid.UUID) error {
	_, err := service.CapaianRepository.DeleteCapaian(capaianID)
	return err
}
