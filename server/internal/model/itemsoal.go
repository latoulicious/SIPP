package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemSoal struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	BankSoal   BankSoal  `gorm:"foreignkey:BankSoalID"` // Foreign key to BankSoal table
	BankSoalID uuid.UUID `gorm:"type:uuid"`
}
