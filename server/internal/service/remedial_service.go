package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type RemedialService struct {
	RemedialRepository *repository.RemedialRepository
}

// NewRemedialService creates a new instance of RemedialService
func NewRemedialService(remedialRepository *repository.RemedialRepository) *RemedialService {
	return &RemedialService{
		RemedialRepository: remedialRepository,
	}
}

// GetRemedial retrieves a list of remedials
func (service *RemedialService) GetRemedial() ([]model.Remedial, error) {
	return service.RemedialRepository.GetRemedial()
}

// GetRemedialByID retrieves a remedial by ID
func (service *RemedialService) GetRemedialByID(remedialID uuid.UUID) (*model.Remedial, error) {
	return service.RemedialRepository.GetRemedialByID(remedialID)
}

// CreateRemedial creates a new remedial
func (service *RemedialService) CreateRemedial(remedial *model.Remedial) error {
	return service.RemedialRepository.CreateRemedial(remedial)
}

// UpdateRemedial updates an existing remedial
func (service *RemedialService) UpdateRemedial(remedial *model.Remedial) error {
	return service.RemedialRepository.UpdateRemedial(remedial)
}

// DeleteRemedial deletes a remedial by ID
func (service *RemedialService) DeleteRemedial(remedialID uuid.UUID) error {
	_, err := service.RemedialRepository.DeleteRemedial(remedialID)
	return err
}
