package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Soal struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID     uuid.UUID `gorm:"type:uuid" json:"userID"`
	User       Users     `gorm:"foreignKey:UserID"`
	MapelID    uuid.UUID `gorm:"type:uuid" json:"mapelId"`
	Mapel      Mapel     `gorm:"foreignKey:MapelID"`
	KelasID    uuid.UUID `gorm:"type:uuid" json:"kelasId"`
	Kelas      Kelas     `gorm:"foreignKey:KelasID"`
	JurusanID  uuid.UUID `gorm:"type:uuid" json:"jurusanId"`
	Jurusan    Jurusan   `gorm:"foreignKey:JurusanID"`
	Hari       string
	Tanggal    string
	Waktu      string
	BankSoalID uuid.UUID `gorm:"type:uuid"`
	BankSoal   BankSoal  `gorm:"foreignkey:BankSoalID"`
	// Not sure if this relations is needed
	// ItemSoal   ItemSoal  `gorm:"foreignkey:ItemSoalID"`
	// ItemSoalID uuid.UUID `gorm:"type:uuid"`
}
