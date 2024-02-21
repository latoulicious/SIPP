package service

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type FormatifService struct {
	FormatifRepository *repository.FormatifRepository
}

// NewFormatifService creates a new instance of FormatifService
func NewFormatifService(formatifRepository *repository.FormatifRepository) *FormatifService {
	return &FormatifService{
		FormatifRepository: formatifRepository,
	}
}

// GetFormatif retrieves a list of formatifs
func (service *FormatifService) GetFormatif() ([]model.Formatif, error) {
	return service.FormatifRepository.GetFormatif()
}

// GetFormatifByID retrieves a formatif by ID
func (service *FormatifService) GetFormatifByID(formatifID uuid.UUID) (*model.Formatif, error) {
	return service.FormatifRepository.GetFormatifByID(formatifID)
}

// CountRelatedQuestionsForAll retrieves the count of related questions for all Formatif records
func (service *FormatifService) CountRelatedQuestionsForAll() ([]*model.Formatif, error) {
	// Log the start of the service method
	log.Println("CountRelatedQuestionsForAll started")

	// Fetch all Formatif records
	var formatifs []*model.Formatif
	if err := service.FormatifRepository.DB.Find(&formatifs).Error; err != nil {
		// Log the error
		log.Printf("Error fetching Formatif records: %v\n", err)
		return nil, err
	}

	// Iterate over each Formatif record
	for _, formatif := range formatifs {
		// Convert DynamicFields to a string and check if it's non-empty
		if string(formatif.DynamicFields) == "" {
			log.Printf("DynamicFields is empty for Formatif ID %s\n", formatif.ID)
			continue
		}

		// Unmarshal the DynamicFields JSON into a Go map
		var dynamicFields map[string]interface{}
		err := json.Unmarshal(formatif.DynamicFields, &dynamicFields)
		if err != nil {
			// Log the error
			log.Printf("Error unmarshaling DynamicFields for Formatif ID %s: %v\n", formatif.ID, err)
			continue
		}

		// Count the number of items in the dynamicFields map
		count := len(dynamicFields)

		// Add the question count to the Formatif object as a new field
		formatif.QuestionCount = int64(count)

	}

	// Log the successful completion of the service method
	log.Println("CountRelatedQuestionsForAll completed successfully")

	return formatifs, nil
}

// CreateFormatif creates a new formatif
func (service *FormatifService) CreateFormatif(formatif *model.Formatif) error {
	return service.FormatifRepository.CreateFormatif(formatif)
}

// UpdateFormatif updates an existing formatif
func (service *FormatifService) UpdateFormatif(formatif *model.Formatif) error {
	return service.FormatifRepository.UpdateFormatif(formatif)
}

// DeleteFormatif deletes a formatif by ID
func (service *FormatifService) DeleteFormatif(formatifID uuid.UUID) error {
	_, err := service.FormatifRepository.DeleteFormatif(formatifID)
	return err
}
