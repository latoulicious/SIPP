package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type SoalService struct {
	SoalRepository *repository.SoalRepository
}

// NewSoalService creates a new instance of SoalService
func NewSoalService(soalRepository *repository.SoalRepository) *SoalService {
	return &SoalService{
		SoalRepository: soalRepository,
	}
}

// GetSoal retrieves a list of soal
func (service *SoalService) GetSoal() ([]model.Soal, error) {
	return service.SoalRepository.GetSoal()
}

// GetSoalByID retrieves a soal by ID
func (service *SoalService) GetSoalByID(soalID uuid.UUID) (*model.Soal, error) {
	return service.SoalRepository.GetSoalByID(soalID)
}

// CreateSoal creates a new soal
func (service *SoalService) CreateSoal(soal *model.Soal) error {
	return service.SoalRepository.CreateSoal(soal)
}

// UpdateSoal updates an existing soal
func (service *SoalService) UpdateSoal(soal *model.Soal) error {
	return service.SoalRepository.UpdateSoal(soal)
}

// DeleteSoal deletes a soal by ID
func (service *SoalService) DeleteSoal(soalID uuid.UUID) error {
	return service.SoalRepository.DeleteSoal(soalID)
}
