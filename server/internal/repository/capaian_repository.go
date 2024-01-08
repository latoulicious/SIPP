package repository

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type CapaianRepository struct {
	DB *gorm.DB
}

func NewCapaianRepository(db *gorm.DB) *CapaianRepository {
	return &CapaianRepository{DB: db}
}

func (repository *CapaianRepository) GetCapaian() ([]model.Capaian, error) {
	var capaian []model.Capaian
	if err := repository.DB.Find(&capaian).Error; err != nil {
		return nil, err
	}
	return capaian, nil
}

func (repository *CapaianRepository) GetCapaianByID(capaianID uuid.UUID) (*model.Capaian, error) {
	var capaian model.Capaian
	if err := repository.DB.First(&capaian, "id = ?", capaianID).Error; err != nil {
		return nil, err
	}
	return &capaian, nil
}

func (repository *CapaianRepository) CreateCapaian(capaian *model.Capaian) error {
	capaian.ID = uuid.New()
	return repository.DB.Create(capaian).Error
}

func (repository *CapaianRepository) UpdateCapaian(capaian *model.Capaian) error {
	return repository.DB.Save(capaian).Error
}

func (repository *CapaianRepository) DeleteCapaian(capaianID uuid.UUID) error {
	// Instead of using Delete method, use Unscoped().Delete to perform a hard delete
	return repository.DB.Unscoped().Delete(&model.Capaian{}, "id = ?", capaianID).Error
}
