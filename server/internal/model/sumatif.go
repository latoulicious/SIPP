package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sumatif struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID      uuid.UUID `gorm:"type:uuid" json:"userID"`
	User        Users     `gorm:"foreignKey:UserID"`
	MapelID     uuid.UUID `gorm:"type:uuid" json:"mapelId"`
	Mapel       Mapel     `gorm:"foreignKey:MapelID"`
	KelasID     uuid.UUID `gorm:"type:uuid" json:"kelasId"`
	Kelas       Kelas     `gorm:"foreignKey:KelasID"`
	TahunAjarID uuid.UUID `gorm:"type:uuid" json:"tahunAjarId"`
	TahunAjar   TahunAjar `gorm:"foreignKey:TahunAjarID"`
	BankSoalID  uuid.UUID `gorm:"type:uuid"`
	BankSoal    BankSoal  `gorm:"foreignkey:BankSoalID"`
}