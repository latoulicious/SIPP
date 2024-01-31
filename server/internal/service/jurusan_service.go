package service

import (
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/repository"
)

type JurusanService struct {
	JurusanRepository *repository.JurusanRepository
}

// NewJurusanService creates a new instance of JurusanService
func NewJurusanService(jurusanRepository *repository.JurusanRepository) *JurusanService {
	return &JurusanService{
		JurusanRepository: jurusanRepository,
	}
}

// GetJurusan retrieves a list of jurusan
func (service *JurusanService) GetJurusan() ([]model.Jurusan, error) {
	return service.JurusanRepository.GetJurusan()
}

// GetJurusanByID retrieves a jurusan by ID
func (service *JurusanService) GetJurusanByID(jurusanID uuid.UUID) (*model.Jurusan, error) {
	return service.JurusanRepository.GetJurusanByID(jurusanID)
}

// GetJurusan Public Endpoint
func (service *JurusanService) GetJurusanPublic() ([]model.Jurusan, error) {
	// Implement logic to fetch all jurusan without requiring JWT authentication
	return service.JurusanRepository.GetJurusanPublic()
}

// CreateJurusan creates a new jurusan
func (service *JurusanService) CreateJurusan(jurusan *model.Jurusan) error {
	return service.JurusanRepository.CreateJurusan(jurusan)
}

// UpdateJurusan updates an existing jurusan
func (service *JurusanService) UpdateJurusan(jurusan *model.Jurusan) error {
	return service.JurusanRepository.UpdateJurusan(jurusan)
}

// DeleteJurusan deletes a jurusan by ID
func (service *JurusanService) DeleteJurusan(jurusanID uuid.UUID) error {
	return service.JurusanRepository.DeleteJurusan(jurusanID)
}
