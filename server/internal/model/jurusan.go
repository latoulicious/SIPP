package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Jurusan struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	Jurusan string    `gorm:"type:varchar(255)"`
}
