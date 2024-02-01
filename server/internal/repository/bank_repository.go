package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type BankRepository struct {
	DB *gorm.DB
}

func NewBankRepository(db *gorm.DB) *BankRepository {
	return &BankRepository{DB: db}
}

func (repository *BankRepository) GetBank() ([]model.BankSoal, error) {
	var bank []model.BankSoal
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Preload("Jurusan").Find(&bank).Error; err != nil {
		return nil, err
	}
	return bank, nil
}

func (repository *BankRepository) GetBankByID(bankID uuid.UUID) (*model.BankSoal, error) {
	var bank model.BankSoal
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Preload("Jurusan").First(&bank, "id = ?", bankID).Error; err != nil {
		return nil, err
	}
	return &bank, nil
}

func (repository *BankRepository) CreateBank(bank *model.BankSoal) error {
	bank.ID = uuid.New()
	if err := repository.DB.Create(bank).Error; err != nil {
		return err
	}
	return repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Preload("Jurusan").First(bank, "id = ?", bank.ID).Error
}

func (repository *BankRepository) UpdateBank(bank *model.BankSoal) error {
	log.Printf("Updating BankSoal with ID: %s\n", bank.ID)

	// Log the values of the related entities
	log.Printf("User: %+v\n", bank.User)
	log.Printf("Mapel: %+v\n", bank.Mapel)
	log.Printf("Kelas: %+v\n", bank.Kelas)
	log.Printf("Jurusan: %+v\n", bank.Jurusan)
	log.Printf("TahunAjar: %+v\n", bank.TahunAjar)

	if err := repository.DB.Save(bank).Error; err != nil {
		log.Printf("Error saving bank: %+v\n", err)
		return err
	}
	return nil
}

func (repository *BankRepository) DeleteBank(bankID uuid.UUID) (*model.BankSoal, error) {
	var bank model.BankSoal
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Preload("Jurusan").First(&bank, "id = ?", bankID).Error; err != nil {
		return nil, err
	}
	if err := repository.DB.Unscoped().Delete(&bank).Error; err != nil {
		return nil, err
	}
	return &bank, nil
}
