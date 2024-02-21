package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type AlurService struct {
	AlurRepository *repository.AlurRepository
}

// NewAlurService creates a new instance of AlurService
func NewAlurService(alurRepository *repository.AlurRepository) *AlurService {
	return &AlurService{
		AlurRepository: alurRepository,
	}
}

// GetAlur retrieves a list of alur
func (service *AlurService) GetAlur() ([]model.AlurTP, error) {
	return service.AlurRepository.GetAlur()
}

// GetAlurByID retrieves an alur by ID
func (service *AlurService) GetAlurByID(alurID uuid.UUID) (*model.AlurTP, error) {
	return service.AlurRepository.GetAlurByID(alurID)
}

// CreateAlur creates a new alur
func (service *AlurService) CreateAlur(alur *model.AlurTP) error {
	return service.AlurRepository.CreateAlur(alur)
}

// UpdateAlur updates an existing alur
func (service *AlurService) UpdateAlur(alur *model.AlurTP) error {
	return service.AlurRepository.UpdateAlur(alur)
}

// DeleteAlur deletes an alur by ID
func (service *AlurService) DeleteAlur(alurID uuid.UUID) error {
	return service.AlurRepository.DeleteAlur(alurID)
}
