package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type CapaianService struct {
	CapaianRepository *repository.CapaianRepository
}

func NewCapaianService(capaianRepository *repository.CapaianRepository) *CapaianService {
	return &CapaianService{
		CapaianRepository: capaianRepository,
	}
}

func (service *CapaianService) GetCapaian() ([]model.Capaian, error) {
	return service.CapaianRepository.GetCapaian()
}

func (service *CapaianService) GetCapaianByID(capaianID uuid.UUID) (*model.Capaian, error) {
	return service.CapaianRepository.GetCapaianByID(capaianID)
}

func (service *CapaianService) CreateCapaian(capaian *model.Capaian) error {
	return service.CapaianRepository.CreateCapaian(capaian)
}

func (service *CapaianService) UpdateCapaian(capaian *model.Capaian) error {
	return service.CapaianRepository.UpdateCapaian(capaian)
}

func (service *CapaianService) DeleteCapaian(capaianID uuid.UUID) error {
	return service.CapaianRepository.DeleteCapaian(capaianID)
}
