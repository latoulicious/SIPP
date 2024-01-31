package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type RemedialRepository struct {
	DB *gorm.DB
}

func NewRemedialRepository(db *gorm.DB) *RemedialRepository {
	return &RemedialRepository{DB: db}
}

func (repository *RemedialRepository) GetRemedial() ([]model.Remedial, error) {
	var remedial []model.Remedial
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").Find(&remedial).Error; err != nil {
		return nil, err
	}
	return remedial, nil
}

func (repository *RemedialRepository) GetRemedialByID(remedialID uuid.UUID) (*model.Remedial, error) {
	var remedial model.Remedial
	if err := repository.DB.Preload("User").Preload("Mapel").Preload("Kelas").Preload("TahunAjar").First(&remedial, "id = ?", remedialID).Error; err != nil {
		return nil, err
	}
	return &remedial, nil
}

func (repository *RemedialRepository) CreateRemedial(remedial *model.Remedial) error {
	// Check if user_id exists in users table
	var user model.Users
	if err := repository.DB.First(&user, "id = ?", remedial.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with id %s does not exist", remedial.UserID)
		}
		return err
	}

	remedial.ID = uuid.New()
	if err := repository.DB.Create(&remedial).Error; err != nil {
		return err
	}
	return nil
}

func (repository *RemedialRepository) UpdateRemedial(remedial *model.Remedial) error {
	log.Printf("Updating Remedial with ID: %s\n", remedial.ID)

	if err := repository.DB.Save(remedial).Error; err != nil {
		log.Printf("Error saving remedial: %+v\n", err)
		return err
	}
	return nil
}

func (repository *RemedialRepository) DeleteRemedial(remedialID uuid.UUID) (*model.Remedial, error) {
	var remedial model.Remedial
	if err := repository.DB.Unscoped().Where("id = ?", remedialID).Delete(&remedial).Error; err != nil {
		return nil, err
	}
	return &remedial, nil
}
