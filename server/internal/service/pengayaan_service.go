package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type PengayaanService struct {
	PengayaanRepository *repository.PengayaanRepository
}

// NewPengayaanService creates a new instance of PengayaanService
func NewPengayaanService(pengayaanRepository *repository.PengayaanRepository) *PengayaanService {
	return &PengayaanService{
		PengayaanRepository: pengayaanRepository,
	}
}

// GetPengayaan retrieves a list of pengayaans
func (service *PengayaanService) GetPengayaan() ([]model.Pengayaan, error) {
	return service.PengayaanRepository.GetPengayaan()
}

// GetPengayaanByID retrieves a pengayaan by ID
func (service *PengayaanService) GetPengayaanByID(pengayaanID uuid.UUID) (*model.Pengayaan, error) {
	return service.PengayaanRepository.GetPengayaanByID(pengayaanID)
}

// CreatePengayaan creates a new pengayaan
func (service *PengayaanService) CreatePengayaan(pengayaan *model.Pengayaan) error {
	return service.PengayaanRepository.CreatePengayaan(pengayaan)
}

// UpdatePengayaan updates an existing pengayaan
func (service *PengayaanService) UpdatePengayaan(pengayaan *model.Pengayaan) error {
	return service.PengayaanRepository.UpdatePengayaan(pengayaan)
}

// DeletePengayaan deletes a pengayaan by ID
func (service *PengayaanService) DeletePengayaan(pengayaanID uuid.UUID) error {
	_, err := service.PengayaanRepository.DeletePengayaan(pengayaanID)
	return err
}
