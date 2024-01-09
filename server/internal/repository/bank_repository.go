package repository

import (
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
	if err := repository.DB.Find(&bank).Error; err != nil {
		return nil, err
	}
	return bank, nil
}

func (repository *BankRepository) GetBankByID(bankID uuid.UUID) (*model.BankSoal, error) {
	var bank model.BankSoal
	if err := repository.DB.First(&bank, "id = ?", bankID).Error; err != nil {
		return nil, err
	}
	return &bank, nil
}

func (repository *BankRepository) CreateBank(bank *model.BankSoal) error {
	bank.ID = uuid.New()
	return repository.DB.Create(bank).Error
}

func (repository *BankRepository) UpdateBank(bank *model.BankSoal) error {
	return repository.DB.Save(bank).Error
}

func (repository *BankRepository) DeleteBank(bankID uuid.UUID) error {
	return repository.DB.Delete(&model.BankSoal{}, "id = ?", bankID).Error
}
