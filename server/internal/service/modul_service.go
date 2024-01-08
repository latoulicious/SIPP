package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type ModulService struct {
	ModulRepository *repository.ModulRepository
}

func NewModulService(modulRepository *repository.ModulRepository) *ModulService {
	return &ModulService{
		ModulRepository: modulRepository,
	}
}

func (service *ModulService) GetModul() ([]model.ModulAjar, error) {
	return service.ModulRepository.GetModul()
}

func (service *ModulService) GetModulByID(modulID uuid.UUID) (model.ModulAjar, error) {
	return service.ModulRepository.GetModulByID(modulID)
}

func (service *ModulService) CreateModul(modul *model.ModulAjar) error {
	return service.ModulRepository.CreateModul(modul)
}

func (service *ModulService) UpdateModul(modul *model.ModulAjar) error {
	return service.ModulRepository.UpdateModul(modul)
}

func (service *ModulService) DeleteModul(modulID uuid.UUID) error {
	return service.ModulRepository.DeleteModul(modulID)
}
