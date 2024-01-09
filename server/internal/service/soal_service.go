package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type SoalService struct {
	SoalRepository *repository.SoalRepository
}

func NewSoalService(soalRepository *repository.SoalRepository) *SoalService {
	return &SoalService{
		SoalRepository: soalRepository,
	}
}

func (service *SoalService) GetSoal() ([]model.Soal, error) {
	return service.SoalRepository.GetSoal()
}

func (service *SoalService) GetSoalByID(soalID uuid.UUID) (*model.Soal, error) {
	return service.SoalRepository.GetSoalByID(soalID)
}

func (service *SoalService) CreateSoal(soal *model.Soal) error {
	return service.SoalRepository.CreateSoal(soal)
}

func (service *SoalService) UpdateSoal(soal *model.Soal) error {
	return service.SoalRepository.UpdateSoal(soal)
}

func (service *SoalService) DeleteSoal(soalID uuid.UUID) error {
	return service.SoalRepository.DeleteSoal(soalID)
}
