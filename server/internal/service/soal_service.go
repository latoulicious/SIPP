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

// GetSoal retrieve a list of soal
func (service *SoalService) GetSoal() ([]*model.Soal, error) {
	var soals []*model.Soal

	// Directly preload dynamic fields
	// Remove the Preload for "Items.DynamicFields" as it is not a valid relation
	err := service.SoalRepository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("Jurusan").Preload("Items").Preload("Items.BankSoal").Find(&soals).Error
	if err != nil {
		return nil, err
	}

	// Count dynamic fields (assuming CountDynamicFields is also updated to work with the new structure)
	for _, soal := range soals {
		for _, item := range soal.Items {
			count, err := service.CountDynamicFields(item.SoalID)
			if err != nil {
				return nil, err
			}
			item.QuestionCount = count
		}
	}

	return soals, nil
}

// GetSoalByID retrieve a soal by id
func (service *SoalService) GetSoalByID(soalID uuid.UUID) (*model.Soal, error) {
	return service.SoalRepository.GetSoalByID(soalID)
}

// CreateSoal creates a new soal
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

// misc function

// Update the CountDynamicFields method in SoalService to use int64
func (service *SoalService) CountDynamicFields(soalID uuid.UUID) (int64, error) {
	var count int64
	err := service.SoalRepository.DB.Model(&model.ItemSoal{}).Where("soal_id = ?", soalID).Count(&count).Error
	return count, err
}
