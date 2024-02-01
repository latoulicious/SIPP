package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pengayaan struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid" json:"userID"`
	User        Users     `gorm:"foreignKey:UserID"`
	MapelID     uuid.UUID `gorm:"type:uuid" json:"mapelId"`
	Mapel       Mapel     `gorm:"foreignKey:MapelID"`
	KelasID     uuid.UUID `gorm:"type:uuid" json:"kelasId"`
	Kelas       Kelas     `gorm:"foreignKey:KelasID"`
	TahunAjarID uuid.UUID `gorm:"type:uuid" json:"tahunAjarId"`
	TahunAjar   TahunAjar `gorm:"foreignKey:TahunAjarID"`
	Pertanyaan  string
}
