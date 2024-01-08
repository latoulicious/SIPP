package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TahunAjar struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;primary_key"`
	Tahun string
}
