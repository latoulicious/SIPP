package repository

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"gorm.io/gorm"
)

type JurusanRepository struct {
	DB *gorm.DB
}

func NewJurusanRepository(db *gorm.DB) *JurusanRepository {
	return &JurusanRepository{DB: db}
}

func (repository *JurusanRepository) GetJurusan() ([]model.Jurusan, error) {
	var jurusan []model.Jurusan
	if err := repository.DB.Find(&jurusan).Error; err != nil {
		return nil, err
	}
	return jurusan, nil
}

func (repository *JurusanRepository) GetJurusanByID(jurusanID uuid.UUID) (*model.Jurusan, error) {
	var jurusan model.Jurusan
	if err := repository.DB.First(&jurusan, "id = ?", jurusanID).Error; err != nil {
		// Log an error if jurusan retrieval fails
		log.Printf("Error fetching jurusan by ID %s: %s\n", jurusanID.String(), err.Error())
		return nil, err
	}

	// Log successful jurusan retrieval
	if os.Getenv("ENVIRONMENT") == "development" {
		log.Printf("Jurusan fetched by ID %s:", jurusanID.String())
	}
	return &jurusan, nil
}

func (repository *JurusanRepository) GetJurusanPublic() ([]model.Jurusan, error) {
	// Implement logic to fetch all jurusan without requiring JWT authentication
	var jurusan []model.Jurusan
	if err := repository.DB.Find(&jurusan).Error; err != nil {
		return nil, err
	}
	return jurusan, nil
}

func (repository *JurusanRepository) CreateJurusan(jurusan *model.Jurusan) error {
	jurusan.ID = uuid.New()
	return repository.DB.Create(jurusan).Error
}

func (repository *JurusanRepository) UpdateJurusan(jurusan *model.Jurusan) error {
	return repository.DB.Save(jurusan).Error
}

func (repository *JurusanRepository) DeleteJurusan(jurusanID uuid.UUID) error {
	// Instead of using Delete method, use Unscoped().Delete to perform a hard delete
	return repository.DB.Unscoped().Delete(&model.Jurusan{}, "id = ?", jurusanID).Error
}
