package service

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type KognitifService struct {
	KognitifRepository *repository.KognitifRepository
}

// NewKognitifService creates a new instance of KognitifService
func NewKognitifService(kognitifRepository *repository.KognitifRepository) *KognitifService {
	return &KognitifService{
		KognitifRepository: kognitifRepository,
	}
}

// GetKognitif retrieves a list of kognitifs
func (service *KognitifService) GetKognitif() ([]model.Kognitif, error) {
	return service.KognitifRepository.GetKognitif()
}

// GetKognitifByID retrieves a kognitif by ID
func (service *KognitifService) GetKognitifByID(kognitifID uuid.UUID) (*model.Kognitif, error) {
	return service.KognitifRepository.GetKognitifByID(kognitifID)
}

// CountRelatedQuestionsForAll retrieves the count of related questions for all Kognitif records
func (service *KognitifService) CountRelatedQuestionsForAll() ([]*model.Kognitif, error) {
	// Log the start of the service method
	log.Println("CountRelatedQuestionsForAll started")

	// Fetch all Kognitif records
	var kognitifs []*model.Kognitif
	if err := service.KognitifRepository.DB.Find(&kognitifs).Error; err != nil {
		// Log the error
		log.Printf("Error fetching Kognitif records: %v\n", err)
		return nil, err
	}

	// Iterate over each Kognitif record
	for _, kognitif := range kognitifs {
		// Convert DynamicFields to a string and check if it's non-empty
		if string(kognitif.DynamicFields) == "" {
			log.Printf("DynamicFields is empty for Kognitif ID %s\n", kognitif.ID)
			continue
		}

		// Unmarshal the DynamicFields JSON into a Go slice
		var dynamicFields []map[string]interface{}
		err := json.Unmarshal(kognitif.DynamicFields, &dynamicFields)
		if err != nil {
			// Log the error
			log.Printf("Error unmarshaling DynamicFields for Kognitif ID %s: %v\n", kognitif.ID, err)
			continue
		}

		// Count the number of items in the dynamicFields slice
		count := len(dynamicFields)

		// Add the question count to the Kognitif object as a new field
		kognitif.QuestionCount = int64(count)
	}

	// Log the successful completion of the service method
	log.Println("CountRelatedQuestionsForAll completed successfully")

	return kognitifs, nil
}

// CreateKognitif creates a new kognitif
func (service *KognitifService) CreateKognitif(kognitif *model.Kognitif) error {
	return service.KognitifRepository.CreateKognitif(kognitif)
}

// UpdateKognitif updates an existing kognitif
func (service *KognitifService) UpdateKognitif(kognitif *model.Kognitif) error {
	return service.KognitifRepository.UpdateKognitif(kognitif)
}

// DeleteKognitif deletes a kognitif by ID
func (service *KognitifService) DeleteKognitif(kognitifID uuid.UUID) error {
	_, err := service.KognitifRepository.DeleteKognitif(kognitifID)
	return err
}
