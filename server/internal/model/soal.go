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
	BankSoal   BankSoal  `gorm:"foreignkey:BankSoalID"` // Foreign key to BankSoal table
	BankSoalID uuid.UUID `gorm:"type:uuid"`
	Hari       string
	Tanggal    string
	Waktu      string
}
