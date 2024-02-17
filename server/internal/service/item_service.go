package service

import (
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type ItemSoalService struct {
	ItemSoalRepository *repository.ItemSoalRepository
}

// NewItemSoalService creates a new instance of ItemSoalService
func NewItemSoalService(itemsoalRepository *repository.ItemSoalRepository) *ItemSoalService {
	return &ItemSoalService{
		ItemSoalRepository: itemsoalRepository,
	}
}

// GetItemSoal retrieves a list of itemsoals
func (service *ItemSoalService) GetItemSoal() ([]model.ItemSoal, error) {
	return service.ItemSoalRepository.GetItemSoal()
}

// GetItemSoalByID retrieves a itemsoal by ID
func (service *ItemSoalService) GetItemSoalByID(itemsoalID uuid.UUID) (*model.ItemSoal, error) {
	return service.ItemSoalRepository.GetItemSoalByID(itemsoalID)
}

// CountRelatedQuestionsForAll retrieves the count of related questions for all ItemSoal records
func (service *ItemSoalService) CountRelatedQuestionsForAll() ([]*model.ItemSoal, error) {
	// Log the start of the service method
	log.Println("CountRelatedQuestionsForAll started")

	// Fetch all ItemSoal records
	var itemsoals []*model.ItemSoal
	if err := service.ItemSoalRepository.DB.Find(&itemsoals).Error; err != nil {
		// Log the error
		log.Printf("Error fetching ItemSoal records: %v\n", err)
		return nil, err
	}

	// Iterate over each ItemSoal record
	for _, itemsoal := range itemsoals {
		// Check if DynamicFields is non-empty
		if len(itemsoal.DynamicFields) == 0 {
			log.Printf("DynamicFields is empty for ItemSoal ID %s\n", itemsoal.ID)
			continue
		}

		// Count the number of items in the DynamicFields map
		count := len(itemsoal.DynamicFields)

		// Add the question count to the ItemSoal object as a new field
		itemsoal.QuestionCount = int64(count)
	}

	// Log the successful completion of the service method
	log.Println("CountRelatedQuestionsForAll completed successfully")

	return itemsoals, nil
}
