package service

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type SumatifService struct {
	SumatifRepository *repository.SumatifRepository
}

// NewSumatifService creates a new instance of SumatifService
func NewSumatifService(sumatifRepository *repository.SumatifRepository) *SumatifService {
	return &SumatifService{
		SumatifRepository: sumatifRepository,
	}
}

// GetSumatif retrieves a list of sumatifs
func (service *SumatifService) GetSumatif() ([]model.Sumatif, error) {
	return service.SumatifRepository.GetSumatif()
}

// GetSumatifByID retrieves a sumatif by ID
func (service *SumatifService) GetSumatifByID(sumatifID uuid.UUID) (*model.Sumatif, error) {
	return service.SumatifRepository.GetSumatifByID(sumatifID)
}

// CountRelatedQuestionsForAll retrieves the count of related questions for all Sumatif records
func (service *SumatifService) CountRelatedQuestionsForAll() ([]*model.Sumatif, error) {
	// Log the start of the service method
	log.Println("CountRelatedQuestionsForAll started")

	// Fetch all Sumatif records
	var sumatifs []*model.Sumatif
	if err := service.SumatifRepository.DB.Find(&sumatifs).Error; err != nil {
		// Log the error
		log.Printf("Error fetching Sumatif records: %v\n", err)
		return nil, err
	}

	// Iterate over each Sumatif record
	for _, sumatif := range sumatifs {
		// Convert DynamicFields to a string and check if it's non-empty
		if string(sumatif.DynamicFields) == "" {
			log.Printf("DynamicFields is empty for Sumatif ID %s\n", sumatif.ID)
			continue
		}

		// Unmarshal the DynamicFields JSON into a Go slice
		var dynamicFields []map[string]interface{}
		err := json.Unmarshal(sumatif.DynamicFields, &dynamicFields)
		if err != nil {
			// Log the error
			log.Printf("Error unmarshaling DynamicFields for Sumatif ID %s: %v\n", sumatif.ID, err)
			continue
		}

		// Count the number of items in the dynamicFields slice
		count := len(dynamicFields)

		// Add the question count to the Sumatif object as a new field
		sumatif.QuestionCount = int64(count)
	}

	// Log the successful completion of the service method
	log.Println("CountRelatedQuestionsForAll completed successfully")

	return sumatifs, nil
}

// CreateSumatif creates a new sumatif
func (service *SumatifService) CreateSumatif(sumatif *model.Sumatif) error {
	return service.SumatifRepository.CreateSumatif(sumatif)
}

// UpdateSumatif updates an existing sumatif
func (service *SumatifService) UpdateSumatif(sumatif *model.Sumatif) error {
	return service.SumatifRepository.UpdateSumatif(sumatif)
}

// DeleteSumatif deletes a sumatif by ID
func (service *SumatifService) DeleteSumatif(sumatifID uuid.UUID) error {
	_, err := service.SumatifRepository.DeleteSumatif(sumatifID)
	return err
}
