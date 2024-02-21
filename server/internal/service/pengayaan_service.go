package service

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type PengayaanService struct {
	PengayaanRepository *repository.PengayaanRepository
}

// NewPengayaanService creates a new instance of PengayaanService
func NewPengayaanService(pengayaanRepository *repository.PengayaanRepository) *PengayaanService {
	return &PengayaanService{
		PengayaanRepository: pengayaanRepository,
	}
}

// GetPengayaan retrieves a list of pengayaans
func (service *PengayaanService) GetPengayaan() ([]model.Pengayaan, error) {
	return service.PengayaanRepository.GetPengayaan()
}

// GetPengayaanByID retrieves a pengayaan by ID
func (service *PengayaanService) GetPengayaanByID(pengayaanID uuid.UUID) (*model.Pengayaan, error) {
	return service.PengayaanRepository.GetPengayaanByID(pengayaanID)
}

// CountRelatedQuestionsForAll retrieves the count of related questions for all Pengayaan records
func (service *PengayaanService) CountRelatedQuestionsForAll() ([]*model.Pengayaan, error) {
	// Log the start of the service method
	log.Println("CountRelatedQuestionsForAll started")

	// Fetch all Pengayaan records
	var pengayaans []*model.Pengayaan
	if err := service.PengayaanRepository.DB.Find(&pengayaans).Error; err != nil {
		// Log the error
		log.Printf("Error fetching Pengayaan records: %v\n", err)
		return nil, err
	}

	// Iterate over each Pengayaan record
	for _, pengayaan := range pengayaans {
		// Convert DynamicFields to a string and check if it's non-empty
		if string(pengayaan.DynamicFields) == "" {
			log.Printf("DynamicFields is empty for Pengayaan ID %s\n", pengayaan.ID)
			continue
		}

		// Unmarshal the DynamicFields JSON into a Go map
		var dynamicFields map[string]interface{}
		err := json.Unmarshal(pengayaan.DynamicFields, &dynamicFields)
		if err != nil {
			// Log the error
			log.Printf("Error unmarshaling DynamicFields for Pengayaan ID %s: %v\n", pengayaan.ID, err)
			continue
		}

		// Count the number of items in the dynamicFields map
		count := len(dynamicFields)

		// Add the question count to the Pengayaan object as a new field
		pengayaan.QuestionCount = int64(count)

	}

	// Log the successful completion of the service method
	log.Println("CountRelatedQuestionsForAll completed successfully")

	return pengayaans, nil
}

// CreatePengayaan creates a new pengayaan
func (service *PengayaanService) CreatePengayaan(pengayaan *model.Pengayaan) error {
	return service.PengayaanRepository.CreatePengayaan(pengayaan)
}

// UpdatePengayaan updates an existing pengayaan
func (service *PengayaanService) UpdatePengayaan(pengayaan *model.Pengayaan) error {
	return service.PengayaanRepository.UpdatePengayaan(pengayaan)
}

// DeletePengayaan deletes a pengayaan by ID
func (service *PengayaanService) DeletePengayaan(pengayaanID uuid.UUID) error {
	_, err := service.PengayaanRepository.DeletePengayaan(pengayaanID)
	return err
}
