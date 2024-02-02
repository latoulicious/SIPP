package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sumatif struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID      uuid.UUID `gorm:"type:uuid" json:"UserID"`
	User        Users     `gorm:"foreignKey:UserID"`
	MapelID     uuid.UUID `gorm:"type:uuid" json:"MapelID"`
	Mapel       Mapel     `gorm:"foreignKey:MapelID"`
	KelasID     uuid.UUID `gorm:"type:uuid" json:"KelasID"`
	Kelas       Kelas     `gorm:"foreignKey:KelasID"`
	TahunAjarID uuid.UUID `gorm:"type:uuid" json:"TahunAjarID"`
	TahunAjar   TahunAjar `gorm:"foreignKey:TahunAjarID"`
	BankSoalID  uuid.UUID `gorm:"type:uuid" json:"BankSoalID"`
	BankSoal    BankSoal  `gorm:"foreignkey:BankSoalID"`
}
