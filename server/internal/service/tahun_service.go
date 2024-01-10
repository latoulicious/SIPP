package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type TahunService struct {
	TahunRepository *repository.TahunRepository
}

// NewTahunService creates a new instance of TahunService
func NewTahunService(tahunRepository *repository.TahunRepository) *TahunService {
	return &TahunService{
		TahunRepository: tahunRepository,
	}
}

// GetTahun retrieves a list of tahun
func (service *TahunService) GetTahun() ([]model.TahunAjar, error) {
	return service.TahunRepository.GetTahun()
}

// GetTahunByID retrieves a tahun by ID
func (service *TahunService) GetTahunByID(tahunID uuid.UUID) (*model.TahunAjar, error) {
	return service.TahunRepository.GetTahunByID(tahunID)
}

// CreateTahun creates a new tahun
func (service *TahunService) CreateTahun(tahun *model.TahunAjar) error {
	return service.TahunRepository.CreateTahun(tahun)
}

// UpdateTahun updates an existing tahun
func (service *TahunService) UpdateTahun(tahun *model.TahunAjar) error {
	return service.TahunRepository.UpdateTahun(tahun)
}

// DeleteTahun deletes a tahun by ID
func (service *TahunService) DeleteTahun(tahunID uuid.UUID) error {
	return service.TahunRepository.DeleteTahun(tahunID)
}
