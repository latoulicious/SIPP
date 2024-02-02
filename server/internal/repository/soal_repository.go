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

// GetSoalWithItems retrieves a Soal along with its associated ItemSoal records
func (repository *SoalRepository) GetSoal(soalID uuid.UUID) (*model.Soal, error) {
	var soal model.Soal
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("Jurusan").Preload("Items").Preload("Items.BankSoal").First(&soal, "id = ?", soalID).Error; err != nil {
		return nil, err
	}
	return &soal, nil
}

func (repository *SoalRepository) GetSoalByID(soalID uuid.UUID) (*model.Soal, error) {
	var soal model.Soal
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("Jurusan").Preload("BankSoal").First(&soal, "id = ?", soalID).Error; err != nil {
		return nil, err
	}
	return &soal, nil
}

// CreateSoalWithItems creates a new Soal and associates it with ItemSoal records
func (repository *SoalRepository) CreateSoal(soal *model.Soal, itemSoalData []*model.ItemSoal) error {
	soal.ID = uuid.New()
	if err := repository.DB.Create(soal).Error; err != nil {
		return err
	}

	// Create ItemSoal records for each BankSoal ID
	for _, itemSoal := range itemSoalData {
		itemSoal.SoalID = soal.ID
		if err := repository.DB.Create(itemSoal).Error; err != nil {
			return err
		}
	}

	return nil
}

func (repository *SoalRepository) UpdateSoal(soal *model.Soal) error {
	// Update the Soal record itself
	if err := repository.DB.Save(soal).Error; err != nil {
		return err
	}

	// No need to update ItemSoal records since they reference the Soal model
	// Any changes to the Soal model will be reflected in the ItemSoal records
	// that reference it due to the foreign key relationship.

	return nil
}

func (repository *SoalRepository) DeleteSoal(soalID uuid.UUID) error {
	// First, delete all associated ItemSoal records
	if err := repository.DB.Where("soal_id = ?", soalID).Delete(&model.ItemSoal{}).Error; err != nil {
		return err
	}

	// Then, delete the Soal record itself
	if err := repository.DB.Delete(&model.Soal{}, "id = ?", soalID).Error; err != nil {
		return err
	}

	return nil
}
