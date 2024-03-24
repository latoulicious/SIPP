package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

// BankService struct that includes all necessary services and repositories
type BankService struct {
	BankRepository *repository.BankRepository
	TingkatService *TingkatService
}

// NewBankService creates a new instance of BankService with all necessary services and repositories
func NewBankService(
	bankRepository *repository.BankRepository,
	tingkatService *TingkatService,
) *BankService {
	return &BankService{
		BankRepository: bankRepository,
		TingkatService: tingkatService,
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

// CreateBank creates a new bank by creating IndikatorTingkat and BankSoal
func (service *BankService) CreateBank(bank *model.BankSoal) error {
	// Create IndikatorTingkat linking the Indikator and Kesukaran
	err := service.TingkatService.CreateIndikatorTingkat(bank.IndikatorID, bank.KesukaranID)
	if err != nil {
		return err
	}

	// Now, create the BankSoal entry with the IDs of the existing Indikator and Kesukaran
	err = service.BankRepository.CreateBank(bank)
	if err != nil {
		return err
	}

	return nil
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
