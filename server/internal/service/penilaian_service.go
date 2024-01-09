package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type PenilaianService struct {
	PenilaianRepository *repository.PenilaianRepository
}

func NewPenilaianService(penilaianRepository *repository.PenilaianRepository) *PenilaianService {
	return &PenilaianService{PenilaianRepository: penilaianRepository}
}

func (service *PenilaianService) GetPenilaian() ([]model.Penilaian, error) {
	return service.PenilaianRepository.GetPenilaian()
}

func (service *PenilaianService) GetPenilaianByID(penilaianID uuid.UUID) (*model.Penilaian, error) {
	return service.PenilaianRepository.GetPenilaianByID(penilaianID)
}

func (service *PenilaianService) CreatePenilaian(penilaian *model.Penilaian) error {
	return service.PenilaianRepository.CreatePenilaian(penilaian)
}

func (service *PenilaianService) UpdatePenilaian(penilaian *model.Penilaian) error {
	return service.PenilaianRepository.UpdatePenilaian(penilaian)
}

func (service *PenilaianService) DeletePenilaian(penilaianID uuid.UUID) error {
	return service.PenilaianRepository.DeletePenilaian(penilaianID)
}
