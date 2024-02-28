package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kelas struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;primary_key"`
	Kelas string    `gorm:"type:varchar(255)"`
}
