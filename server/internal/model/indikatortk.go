package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IndikatorTingkat struct {
	gorm.Model
	ID          uuid.UUID  `gorm:"type:uuid;primary_key"`
	IndikatorID uuid.UUID  `gorm:"type:uuid"`
	KesukaranID uuid.UUID  `gorm:"type:uuid"`
	Indikator   Indikator  `gorm:"foreignKey:IndikatorID"`
	Kesukaran   Kesukaran  `gorm:"foreignKey:KesukaranID"`
	BankSoal    []BankSoal `gorm:"foreignKey:IndikatorTingkatID"`
}
