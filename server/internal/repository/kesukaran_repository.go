package repository

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type KesukaranRepository struct {
	DB *gorm.DB
}

func NewKesukaranRepository(db *gorm.DB) *KesukaranRepository {
	return &KesukaranRepository{DB: db}
}

func (repository *KesukaranRepository) GetKesukaran() ([]model.Kesukaran, error) {
	var kesukaran []model.Kesukaran
	if err := repository.DB.Find(&kesukaran).Error; err != nil {
		return nil, err
	}

	return kesukaran, nil
}

func (repository *KesukaranRepository) GetKesukaranByID(kesukaranID uuid.UUID) (*model.Kesukaran, error) {
	var kesukaran model.Kesukaran
	if err := repository.DB.First(&kesukaran, "id =?", kesukaranID).Error; err != nil {
		return nil, err
	}
	return &kesukaran, nil
}

func (repository *KesukaranRepository) CreateKesukaran(kesukaran *model.Kesukaran) error {
	kesukaran.ID = uuid.New()
	return repository.DB.Create(kesukaran).Error
}

func (repository *KesukaranRepository) UpdateKesukaran(kesukaran *model.Kesukaran) error {
	return repository.DB.Save(kesukaran).Error
}

func (repository *KesukaranRepository) DeleteKesukaran(kesukaranID uuid.UUID) error {
	if err := repository.DB.Unscoped().Where("id =?", kesukaranID).Delete(&model.Kesukaran{}).Error; err != nil {
		return err
	}

	return nil
}
