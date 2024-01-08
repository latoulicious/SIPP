package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Soal struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	Kelas      Kelas     `gorm:"foreignkey:KelasID"` // Foreign key to Kelas table
	KelasID    uuid.UUID `gorm:"type:uuid"`
	Hari       string
	Tanggal    string
	Waktu      string
	ItemSoal   ItemSoal  `gorm:"foreignkey:ItemSoalID"`
	ItemSoalID uuid.UUID `gorm:"type:uuid"`
}
