package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type SumatifService struct {
	SumatifRepository *repository.SumatifRepository
}

// NewSumatifService creates a new instance of SumatifService
func NewSumatifService(sumatifRepository *repository.SumatifRepository) *SumatifService {
	return &SumatifService{
		SumatifRepository: sumatifRepository,
	}
}

// GetSumatif retrieves a list of sumatifs
func (service *SumatifService) GetSumatif() ([]model.Sumatif, error) {
	return service.SumatifRepository.GetSumatif()
}

// GetSumatifByID retrieves a sumatif by ID
func (service *SumatifService) GetSumatifByID(sumatifID uuid.UUID) (*model.Sumatif, error) {
	return service.SumatifRepository.GetSumatifByID(sumatifID)
}

// CreateSumatif creates a new sumatif
func (service *SumatifService) CreateSumatif(sumatif *model.Sumatif) error {
	return service.SumatifRepository.CreateSumatif(sumatif)
}

// UpdateSumatif updates an existing sumatif
func (service *SumatifService) UpdateSumatif(sumatif *model.Sumatif) error {
	return service.SumatifRepository.UpdateSumatif(sumatif)
}

// DeleteSumatif deletes a sumatif by ID
func (service *SumatifService) DeleteSumatif(sumatifID uuid.UUID) error {
	_, err := service.SumatifRepository.DeleteSumatif(sumatifID)
	return err
}
