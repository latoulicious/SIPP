package repository

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type AlurRepository struct {
	DB *gorm.DB
}

func NewAlurRepository(db *gorm.DB) *AlurRepository {
	return &AlurRepository{DB: db}
}

func (repository *AlurRepository) GetAlur() ([]model.AlurTP, error) {
	var alur []model.AlurTP
	if err := repository.DB.Find(&alur).Error; err != nil {
		return nil, err
	}
	return alur, nil
}

func (repository *AlurRepository) GetAlurByID(alurID uuid.UUID) (*model.AlurTP, error) {
	var alur model.AlurTP
	if err := repository.DB.First(&alur, "id = ?", alurID).Error; err != nil {
		return nil, err
	}
	return &alur, nil
}

func (repository *AlurRepository) CreateAlur(alur *model.AlurTP) error {
	alur.ID = uuid.New()
	return repository.DB.Create(alur).Error
}

func (repository *AlurRepository) UpdateAlur(alur *model.AlurTP) error {
	return repository.DB.Save(alur).Error
}

func (repository *AlurRepository) DeleteAlur(alurID uuid.UUID) error {
	// Instead of using Delete method, use Unscoped().Delete to perform a hard delete
	return repository.DB.Unscoped().Delete(&model.AlurTP{}, "id = ?", alurID).Error
}
