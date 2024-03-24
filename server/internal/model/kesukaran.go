package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kesukaran struct {
	gorm.Model
	ID               uuid.UUID          `gorm:"type:uuid;primary_key"`
	TingkatKesukaran string             `gorm:"type:varchar(255)"`
	IndikatorTingkat []IndikatorTingkat `gorm:"foreignKey:KesukaranID"`
}
