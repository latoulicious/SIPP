package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemSoal struct {
	gorm.Model
	ID            uuid.UUID              `gorm:"type:uuid;primary_key"`
	SoalID        uuid.UUID              `gorm:"type:uuid"`
	Soal          *Soal                  `gorm:"foreignKey:SoalID;associationForeignKey:SoalID"`
	BankSoalID    uuid.UUID              `gorm:"type:uuid"`
	BankSoal      BankSoal               `gorm:"foreignKey:BankSoalID"`
	DynamicFields map[string]interface{} `gorm:"type:jsonb"`
}
