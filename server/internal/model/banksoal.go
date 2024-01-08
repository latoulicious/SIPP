package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankSoal struct {
	gorm.Model
	ID               uuid.UUID `gorm:"type:uuid;primary_key"`
	Soal             string
	OptionA          string
	OptionB          string
	OptionC          string
	OptionD          string
	OptionE          string
	KunciJawaban     string
	Materi           string
	Indikator        string
	TingkatKesukaran string
	FaktorKognitif   string
}
