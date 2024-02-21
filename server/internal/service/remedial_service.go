package service

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type RemedialService struct {
	RemedialRepository *repository.RemedialRepository
}

// NewRemedialService creates a new instance of RemedialService
func NewRemedialService(remedialRepository *repository.RemedialRepository) *RemedialService {
	return &RemedialService{
		RemedialRepository: remedialRepository,
	}
}

// GetRemedial retrieves a list of remedials
func (service *RemedialService) GetRemedial() ([]model.Remedial, error) {
	return service.RemedialRepository.GetRemedial()
}

// GetRemedialByID retrieves a remedial by ID
func (service *RemedialService) GetRemedialByID(remedialID uuid.UUID) (*model.Remedial, error) {
	return service.RemedialRepository.GetRemedialByID(remedialID)
}

// CountRelatedQuestionsForAll retrieves the count of related questions for all Remedial records
func (service *RemedialService) CountRelatedQuestionsForAll() ([]*model.Remedial, error) {
	// Log the start of the service method
	log.Println("CountRelatedQuestionsForAll started")

	// Fetch all Remedial records
	var remedials []*model.Remedial
	if err := service.RemedialRepository.DB.Find(&remedials).Error; err != nil {
		// Log the error
		log.Printf("Error fetching Remedial records: %v\n", err)
		return nil, err
	}

	// Iterate over each Remedial record
	for _, remedial := range remedials {
		// Convert DynamicFields to a string and check if it's non-empty
		if string(remedial.DynamicFields) == "" {
			log.Printf("DynamicFields is empty for Remedial ID %s\n", remedial.ID)
			continue
		}

		// Unmarshal the DynamicFields JSON into a Go map
		var dynamicFields map[string]interface{}
		err := json.Unmarshal(remedial.DynamicFields, &dynamicFields)
		if err != nil {
			// Log the error
			log.Printf("Error unmarshaling DynamicFields for Remedial ID %s: %v\n", remedial.ID, err)
			continue
		}

		// Count the number of items in the dynamicFields map
		count := len(dynamicFields)

		// Add the question count to the Remedial object as a new field
		remedial.QuestionCount = int64(count)

	}

	// Log the successful completion of the service method
	log.Println("CountRelatedQuestionsForAll completed successfully")

	return remedials, nil
}

// CreateRemedial creates a new remedial
func (service *RemedialService) CreateRemedial(remedial *model.Remedial) error {
	return service.RemedialRepository.CreateRemedial(remedial)
}

// UpdateRemedial updates an existing remedial
func (service *RemedialService) UpdateRemedial(remedial *model.Remedial) error {
	return service.RemedialRepository.UpdateRemedial(remedial)
}

// DeleteRemedial deletes a remedial by ID
func (service *RemedialService) DeleteRemedial(remedialID uuid.UUID) error {
	_, err := service.RemedialRepository.DeleteRemedial(remedialID)
	return err
}
