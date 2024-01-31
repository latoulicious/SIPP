package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type KognitifService struct {
	KognitifRepository *repository.KognitifRepository
}

// NewKognitifService creates a new instance of KognitifService
func NewKognitifService(kognitifRepository *repository.KognitifRepository) *KognitifService {
	return &KognitifService{
		KognitifRepository: kognitifRepository,
	}
}

// GetKognitif retrieves a list of kognitifs
func (service *KognitifService) GetKognitif() ([]model.Kognitif, error) {
	return service.KognitifRepository.GetKognitif()
}

// GetKognitifByID retrieves a kognitif by ID
func (service *KognitifService) GetKognitifByID(kognitifID uuid.UUID) (*model.Kognitif, error) {
	return service.KognitifRepository.GetKognitifByID(kognitifID)
}

// CreateKognitif creates a new kognitif
func (service *KognitifService) CreateKognitif(kognitif *model.Kognitif) error {
	return service.KognitifRepository.CreateKognitif(kognitif)
}

// UpdateKognitif updates an existing kognitif
func (service *KognitifService) UpdateKognitif(kognitif *model.Kognitif) error {
	return service.KognitifRepository.UpdateKognitif(kognitif)
}

// DeleteKognitif deletes a kognitif by ID
func (service *KognitifService) DeleteKognitif(kognitifID uuid.UUID) error {
	_, err := service.KognitifRepository.DeleteKognitif(kognitifID)
	return err
}
