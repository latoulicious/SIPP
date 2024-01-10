package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type ModulService struct {
	ModulRepository *repository.ModulRepository
}

// NewModulService creates a new instance of ModulService
func NewModulService(modulRepository *repository.ModulRepository) *ModulService {
	return &ModulService{
		ModulRepository: modulRepository,
	}
}

// GetModul retrieves a list of moduls
func (service *ModulService) GetModul() ([]model.ModulAjar, error) {
	return service.ModulRepository.GetModul()
}

// GetModulByID retrieves a modul by ID
func (service *ModulService) GetModulByID(modulID uuid.UUID) (model.ModulAjar, error) {
	return service.ModulRepository.GetModulByID(modulID)
}

// CreateModul creates a new modul
func (service *ModulService) CreateModul(modul *model.ModulAjar) error {
	return service.ModulRepository.CreateModul(modul)
}

// UpdateModul updates an existing modul
func (service *ModulService) UpdateModul(modul *model.ModulAjar) error {
	return service.ModulRepository.UpdateModul(modul)
}

// DeleteModul deletes a modul by ID
func (service *ModulService) DeleteModul(modulID uuid.UUID) error {
	return service.ModulRepository.DeleteModul(modulID)
}
