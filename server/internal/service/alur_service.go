package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type AlurService struct {
	AlurRepository *repository.AlurRepository
}

func NewAlurService(alurRepository *repository.AlurRepository) *AlurService {
	return &AlurService{
		AlurRepository: alurRepository,
	}
}

func (service *AlurService) GetAlur() ([]model.AlurTP, error) {
	return service.AlurRepository.GetAlur()
}

func (service *AlurService) GetAlurByID(alurID uuid.UUID) (*model.AlurTP, error) {
	return service.AlurRepository.GetAlurByID(alurID)
}

func (service *AlurService) CreateAlur(alur *model.AlurTP) error {
	return service.AlurRepository.CreateAlur(alur)
}

func (service *AlurService) UpdateAlur(alur *model.AlurTP) error {
	return service.AlurRepository.UpdateAlur(alur)
}

func (service *AlurService) DeleteAlur(alurID uuid.UUID) error {
	return service.AlurRepository.DeleteAlur(alurID)
}
