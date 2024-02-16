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

// GetSoal retrieves all Soal records along with their associated ItemSoal records
func (service *SoalService) GetSoal() ([]*model.Soal, error) {
	return service.SoalRepository.GetSoal()
}

// GetSoalByID retrieve a soal by id
func (service *SoalService) GetSoalByID(soalID uuid.UUID) (*model.Soal, error) {
	return service.SoalRepository.GetSoalByID(soalID)
}

func (service *SoalService) CreateSoal(soal *model.Soal, itemSoal *model.ItemSoal) error {
	// Save the Soal record
	if err := service.SoalRepository.CreateSoal(soal); err != nil {
		return err
	}

	// Associate the ItemSoal with the newly created Soal
	itemSoal.SoalID = soal.ID

	// Save the ItemSoal record
	if err := service.SoalRepository.CreateItem(itemSoal); err != nil {
		return err
	}

	return nil // Successful creation
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
