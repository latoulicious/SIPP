package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type PenilaianService struct {
	PenilaianRepository *repository.PenilaianRepository
}

// NewPenilaianService creates a new instance of PenilaianService
func NewPenilaianService(penilaianRepository *repository.PenilaianRepository) *PenilaianService {
	return &PenilaianService{
		PenilaianRepository: penilaianRepository,
	}
}

// GetPenilaian retrieves a list of penilaian
func (service *PenilaianService) GetPenilaian() ([]model.Penilaian, error) {
	return service.PenilaianRepository.GetPenilaian()
}

// GetPenilaianByID retrieves a penilaian by ID
func (service *PenilaianService) GetPenilaianByID(penilaianID uuid.UUID) (*model.Penilaian, error) {
	return service.PenilaianRepository.GetPenilaianByID(penilaianID)
}

// CreatePenilaian creates a new penilaian
func (service *PenilaianService) CreatePenilaian(penilaian *model.Penilaian) error {
	return service.PenilaianRepository.CreatePenilaian(penilaian)
}

// UpdatePenilaian updates an existing penilaian
func (service *PenilaianService) UpdatePenilaian(penilaian *model.Penilaian) error {
	return service.PenilaianRepository.UpdatePenilaian(penilaian)
}

// DeletePenilaian deletes a penilaian by ID
func (service *PenilaianService) DeletePenilaian(penilaianID uuid.UUID) error {
	return service.PenilaianRepository.DeletePenilaian(penilaianID)
}
