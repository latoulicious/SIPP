package model

import (
	"github.com/google/uuid"
)

type Kognitif struct {
	BaseModel
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"` // New ID field for Kognitif
	BankSoalID    uuid.UUID `gorm:"type:uuid" json:"BankSoalID"`
	BankSoal      BankSoal  `gorm:"foreignkey:BankSoalID"`
	QuestionCount int64     `json:"questionCount"`
}
