package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type TahunService struct {
	TahunRepository *repository.TahunRepository
}

func NewTahunService(tahunRepository *repository.TahunRepository) *TahunService {
	return &TahunService{
		TahunRepository: tahunRepository,
	}
}

func (service *TahunService) GetTahun() ([]model.TahunAjar, error) {
	return service.TahunRepository.GetTahun()
}

func (service *TahunService) GetTahunByID(tahunID uuid.UUID) (*model.TahunAjar, error) {
	return service.TahunRepository.GetTahunByID(tahunID)
}

func (service *TahunService) CreateTahun(tahun *model.TahunAjar) error {
	return service.TahunRepository.CreateTahun(tahun)
}

func (service *TahunService) UpdateTahun(tahun *model.TahunAjar) error {
	return service.TahunRepository.UpdateTahun(tahun)
}

func (service *TahunService) DeleteTahun(tahunID uuid.UUID) error {
	return service.TahunRepository.DeleteTahun(tahunID)
}
