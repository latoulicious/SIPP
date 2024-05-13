package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type IndikatorService struct {
	IndikatorRepository *repository.IndikatorRepository
}

// NewIndikatorService creates a new instance of IndikatorService
func NewIndikatorService(indikatorRepository *repository.IndikatorRepository) *IndikatorService {
	return &IndikatorService{
		IndikatorRepository: indikatorRepository,
	}
}

// GetIndikator retrieve a list of indikator
func (service *IndikatorService) GetIndikator() ([]model.Indikator, error) {
	return service.IndikatorRepository.GetIndikator()
}

// GetIndikatorByID retrieve a indikator by id
func (service *IndikatorService) GetIndikatorByID(indikatorID uuid.UUID) (*model.Indikator, error) {
	return service.IndikatorRepository.GetIndikatorByID(indikatorID)
}

// CreateIndikator creates a new indikator
func (service *IndikatorService) CreateIndikator(indikator *model.Indikator) error {
	return service.IndikatorRepository.CreateIndikator(indikator)
}

// UpdateIndikator updates an existing indikator
func (service *IndikatorService) UpdateIndikator(indikator *model.Indikator) error {
	return service.IndikatorRepository.UpdateIndikator(indikator)
}

// DeleteIndikator deletes a indikator by ID
func (service *IndikatorService) DeleteIndikator(indikatorID uuid.UUID) error {
	return service.IndikatorRepository.DeleteIndikator(indikatorID)
}
