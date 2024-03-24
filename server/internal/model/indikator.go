package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Indikator struct {
	gorm.Model
	ID               uuid.UUID          `gorm:"type:uuid;primary_key"`
	Materi           string             `gorm:"type:varchar(255)"`
	Indikator        string             `gorm:"type:varchar(255)"`
	IndikatorTingkat []IndikatorTingkat `gorm:"foreignKey:IndikatorID"`
}
