package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type FormatifService struct {
	FormatifRepository *repository.FormatifRepository
}

// NewFormatifService creates a new instance of FormatifService
func NewFormatifService(formatifRepository *repository.FormatifRepository) *FormatifService {
	return &FormatifService{
		FormatifRepository: formatifRepository,
	}
}

// GetFormatif retrieves a list of formatifs
func (service *FormatifService) GetFormatif() ([]model.Formatif, error) {
	return service.FormatifRepository.GetFormatif()
}

// GetFormatifByID retrieves a formatif by ID
func (service *FormatifService) GetFormatifByID(formatifID uuid.UUID) (*model.Formatif, error) {
	return service.FormatifRepository.GetFormatifByID(formatifID)
}

// CreateFormatif creates a new formatif
func (service *FormatifService) CreateFormatif(formatif *model.Formatif) error {
	return service.FormatifRepository.CreateFormatif(formatif)
}

// UpdateFormatif updates an existing formatif
func (service *FormatifService) UpdateFormatif(formatif *model.Formatif) error {
	return service.FormatifRepository.UpdateFormatif(formatif)
}

// DeleteFormatif deletes a formatif by ID
func (service *FormatifService) DeleteFormatif(formatifID uuid.UUID) error {
	_, err := service.FormatifRepository.DeleteFormatif(formatifID)
	return err
}
