package repository

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type SoalRepository struct {
	DB *gorm.DB
}

func NewSoalRepository(db *gorm.DB) *SoalRepository {
	return &SoalRepository{DB: db}
}

// GetSoal retrieves all Soal records along with their associated ItemSoal records
func (repository *SoalRepository) GetSoal() ([]*model.Soal, error) {
	var soals []*model.Soal
	err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("Jurusan").Find(&soals).Error
	if err != nil {
		return nil, err
	}
	return soals, nil
}

// GetSoalByID retrieves a Soal along with its associated ItemSoal records by ID
func (repository *SoalRepository) GetSoalByID(soalID uuid.UUID) (*model.Soal, error) {
	var soal model.Soal
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("Jurusan").Preload("Items").Preload("Items.BankSoal").Preload("Indikator").Preload("Kesukaran").First(&soal, "id = ?", soalID).Error; err != nil {
		return nil, err
	}
	return &soal, nil
}

// GetSoalByIndikatorTingkatID retrieves all Soal records associated with a specific IndikatorTingkat ID
func (repository *SoalRepository) GetSoalByIndikatorTingkatID(indikatorTingkatID uuid.UUID) ([]*model.Soal, error) {
	var soals []*model.Soal
	err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("Jurusan").Preload("Items").Preload("Items.BankSoal").Where("indikator_tingkat_id = ?", indikatorTingkatID).Find(&soals).Error
	if err != nil {
		return nil, err
	}
	return soals, nil
}

func (repository *SoalRepository) CreateSoal(soal *model.Soal) error {
	soal.ID = uuid.New()
	return repository.DB.Create(soal).Error
}

func (repository *SoalRepository) CreateItem(item *model.ItemSoal) error {
	item.ID = uuid.New()
	return repository.DB.Create(item).Error
}

func (repository *SoalRepository) UpdateSoal(soal *model.Soal) error {
	// Update the Soal record itself
	if err := repository.DB.Save(soal).Error; err != nil {
		return err
	}
	return nil
}

// DeleteSoal deletes a soal by ID and its associated ItemSoal records
func (repository *SoalRepository) DeleteSoal(soalID uuid.UUID) error {
	// First, hard delete all associated ItemSoal records
	if err := repository.DB.Unscoped().Where("soal_id = ?", soalID).Delete(&model.ItemSoal{}).Error; err != nil {
		return err
	}

	// Then, hard delete the Soal record itself
	if err := repository.DB.Unscoped().Delete(&model.Soal{}, "id = ?", soalID).Error; err != nil {
		return err
	}

	return nil
}
