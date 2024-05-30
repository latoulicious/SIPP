package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mapel struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;primary_key"`
	Mapel string    `gorm:"type:varchar(50)"`
}
