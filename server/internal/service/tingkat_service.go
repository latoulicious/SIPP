package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type TingkatService struct {
	TingkatRepository *repository.TingkatRepository
}

// NewTingkatService creates a new instance of TingkatService
func NewTingkatService(tingkatRepository *repository.TingkatRepository) *TingkatService {
	return &TingkatService{
		TingkatRepository: tingkatRepository,
	}
}

// GetTingkat retrieve a list of tingkat
func (service *TingkatService) GetIndikatorTingkat() ([]model.IndikatorTingkat, error) {
	return service.TingkatRepository.GetIndikatorTingkat()
}

// GetTingkatByID retrieve a tingkat by id
func (service *TingkatService) GetIndikatorTingkatByID(id uuid.UUID) (*model.IndikatorTingkat, error) {
	return service.TingkatRepository.GetIndikatorTingkatByID(id)
}

// CreateTingkat creates a new tingkat
func (service *TingkatService) CreateIndikatorTingkat(indikatorID uuid.UUID, kesukaranID uuid.UUID) error {
	return service.TingkatRepository.CreateIndikatorTingkat(indikatorID, kesukaranID)
}
