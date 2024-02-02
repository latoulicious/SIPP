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

// GetSoalWithItems retrieves a Soal along with its associated ItemSoal records
func (service *SoalService) GetSoal(soalID uuid.UUID) (*model.Soal, error) {
	return service.SoalRepository.GetSoal(soalID)
}

// GetSoalByID retrieve a soal by id
func (service *SoalService) GetSoalByID(soalID uuid.UUID) (*model.Soal, error) {
	return service.SoalRepository.GetSoalByID(soalID)
}

// CreateSoalWithItems creates a new Soal and associates it with ItemSoal records
func (service *SoalService) CreateSoal(soal *model.Soal, itemSoalData []*model.ItemSoal) error {
	return service.SoalRepository.CreateSoal(soal, itemSoalData)
}

// UpdateSoal updates an existing soal
func (service *SoalService) UpdateSoal(soal *model.Soal) error {
	return service.SoalRepository.UpdateSoal(soal)
}

// DeleteSoal deletes a soal by ID
func (service *SoalService) DeleteSoal(soalID uuid.UUID) error {
	err := service.SoalRepository.DeleteSoal(soalID)
	return err
}
