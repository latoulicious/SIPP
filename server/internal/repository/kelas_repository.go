package repository

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type KelasRepository struct {
	DB *gorm.DB
}

func NewKelasRepository(db *gorm.DB) *KelasRepository {
	return &KelasRepository{DB: db}
}

func (repository *KelasRepository) GetKelas() ([]model.Kelas, error) {
	var kelas []model.Kelas
	if err := repository.DB.Find(&kelas).Error; err != nil {
		return nil, err
	}
	return kelas, nil
}

func (repository *KelasRepository) GetKelasByID(kelasID uuid.UUID) (*model.Kelas, error) {
	var kelas model.Kelas
	if err := repository.DB.First(&kelas, "id = ?", kelasID).Error; err != nil {
		// Log an error if kelas retrieval fails
		log.Printf("Error fetching kelas by ID %s: %s\n", kelasID.String(), err.Error())
		return nil, err
	}

	// Log successful kelas retrieval
	if os.Getenv("ENVIRONMENT") == "development" {
		log.Printf("Kelas fetched by ID %s:", kelasID.String())
	}
	return &kelas, nil
}

func (repository *KelasRepository) GetKelasPublic() ([]model.Kelas, error) {
	// Implement logic to fetch all kelas without requiring JWT authentication
	var kelas []model.Kelas
	if err := repository.DB.Find(&kelas).Error; err != nil {
		return nil, err
	}
	return kelas, nil
}

func (repository *KelasRepository) CreateKelas(kelas *model.Kelas) error {
	kelas.ID = uuid.New()
	return repository.DB.Create(kelas).Error
}

func (repository *KelasRepository) UpdateKelas(kelas *model.Kelas) error {
	return repository.DB.Save(kelas).Error
}

func (repository *KelasRepository) DeleteKelas(kelasID uuid.UUID) error {
	// Instead of using Delete method, use Unscoped().Delete to perform a hard delete
	return repository.DB.Unscoped().Delete(&model.Kelas{}, "id = ?", kelasID).Error
}
