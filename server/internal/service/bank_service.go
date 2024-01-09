package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type BankService struct {
	BankRepository *repository.BankRepository
}

func NewBankService(bankRepository *repository.BankRepository) *BankService {
	return &BankService{
		BankRepository: bankRepository,
	}
}

func (service *BankService) GetBank() ([]model.BankSoal, error) {
	return service.BankRepository.GetBank()
}

func (service *BankService) GetBankByID(bankID uuid.UUID) (*model.BankSoal, error) {
	return service.BankRepository.GetBankByID(bankID)
}

func (service *BankService) CreateBank(bank *model.BankSoal) error {
	return service.BankRepository.CreateBank(bank)
}

func (service *BankService) UpdateBank(bank *model.BankSoal) error {
	return service.BankRepository.UpdateBank(bank)
}

func (service *BankService) DeleteBank(bankID uuid.UUID) error {
	return service.BankRepository.DeleteBank(bankID)
}
