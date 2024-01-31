package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type BankService struct {
	BankRepository *repository.BankRepository
}

// NewBankService creates a new instance of BankService
func NewBankService(bankRepository *repository.BankRepository) *BankService {
	return &BankService{
		BankRepository: bankRepository,
	}
}

// GetBank retrieve a list of bank
func (service *BankService) GetBank() ([]model.BankSoal, error) {
	return service.BankRepository.GetBank()
}

// GetBankByID retrieve a bank by id
func (service *BankService) GetBankByID(bankID uuid.UUID) (*model.BankSoal, error) {
	return service.BankRepository.GetBankByID(bankID)
}

// CreateBank creates a new bank
func (service *BankService) CreateBank(bank *model.BankSoal) error {
	return service.BankRepository.CreateBank(bank)
}

// UpdateBank updates an existing bank
func (service *BankService) UpdateBank(bank *model.BankSoal) error {
	return service.BankRepository.UpdateBank(bank)
}

// DeleteBank deletes a bank by ID
func (service *BankService) DeleteBank(bankID uuid.UUID) error {
	_, err := service.BankRepository.DeleteBank(bankID)
	return err
}
