package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type KelasService struct {
	KelasRepository *repository.KelasRepository
}

// NewKelasService creates a new instance of KelasService
func NewKelasService(kelasRepository *repository.KelasRepository) *KelasService {
	return &KelasService{
		KelasRepository: kelasRepository,
	}
}

// GetKelas retrieves a list of kelas
func (service *KelasService) GetKelas() ([]model.Kelas, error) {
	return service.KelasRepository.GetKelas()
}

// GetKelasByID retrieves a kelas by ID
func (service *KelasService) GetKelasByID(kelasID uuid.UUID) (*model.Kelas, error) {
	return service.KelasRepository.GetKelasByID(kelasID)
}

// CreateKelas creates a new kelas
func (service *KelasService) CreateKelas(kelas *model.Kelas) error {
	return service.KelasRepository.CreateKelas(kelas)
}

// UpdateKelas updates an existing kelas
func (service *KelasService) UpdateKelas(kelas *model.Kelas) error {
	return service.KelasRepository.UpdateKelas(kelas)
}

// DeleteKelas deletes a kelas by ID
func (service *KelasService) DeleteKelas(kelasID uuid.UUID) error {
	return service.KelasRepository.DeleteKelas(kelasID)
}
