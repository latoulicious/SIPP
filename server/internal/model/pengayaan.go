package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pengayaan struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid" json:"UserID"`
	User        Users     `gorm:"foreignKey:UserID"`
	MapelID     uuid.UUID `gorm:"type:uuid" json:"MapelID"`
	Mapel       Mapel     `gorm:"foreignKey:MapelID"`
	KelasID     uuid.UUID `gorm:"type:uuid" json:"KelasID"`
	Kelas       Kelas     `gorm:"foreignKey:KelasID"`
	TahunAjarID uuid.UUID `gorm:"type:uuid" json:"TahunAjarID"`
	TahunAjar   TahunAjar `gorm:"foreignKey:TahunAjarID"`
	Pertanyaan  string    `gorm:"type:varchar(255)"`
}
