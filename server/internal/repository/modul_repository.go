package repository

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type ModulRepository struct {
	DB *gorm.DB
}

func NewModulRepository(db *gorm.DB) *ModulRepository {
	return &ModulRepository{DB: db}
}

func (repository *ModulRepository) GetModul() ([]model.ModulAjar, error) {
	var modul []model.ModulAjar
	if err := repository.DB.Find(&modul).Error; err != nil {
		return nil, err
	}
	return modul, nil
}

func (repository *ModulRepository) GetModulByID(modulID uuid.UUID) (model.ModulAjar, error) {
	var modul model.ModulAjar
	if err := repository.DB.First(&modul, "id = ?", modulID).Error; err != nil {
		return model.ModulAjar{}, err
	}
	return modul, nil
}

func (repository *ModulRepository) CreateModul(modul *model.ModulAjar) error {
	modul.ID = uuid.New()
	return repository.DB.Create(modul).Error
}

func (repository *ModulRepository) UpdateModul(modul *model.ModulAjar) error {
	return repository.DB.Save(modul).Error
}

func (repository *ModulRepository) DeleteModul(modulID uuid.UUID) error {
	// Instead of using Delete method, use Unscoped().Delete to perform a hard delete
	return repository.DB.Delete(&model.ModulAjar{}, "id = ?", modulID).Error
}
